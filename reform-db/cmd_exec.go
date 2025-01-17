package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/qw4n7y/reform"
)

var (
	execFlags  = flag.NewFlagSet("exec", flag.ExitOnError)
	execSplitF = execFlags.Bool("split", false, "Split statements by semicolon; does not handles them in string literals")
)

func init() {
	execFlags.Usage = func() {
		fmt.Fprintf(os.Stderr, "`exec` command executes SQL queries from given files or stdin.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [global flags] exec [exec flags] [file names]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Global flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExec flags:\n")
		execFlags.PrintDefaults()

		// TODO mention -split flag
		fmt.Fprintf(os.Stderr, `
Each file's content is executed as a single query. If it contains multiple
statements, make sure SQL driver supports them. If file names are not given,
a query is read from stdin until EOF, then executed.
`)
	}
}

// readFiles reads queries from given files, or from stdin, if files are not given
func readFiles(files []string, split bool) (queries []string) {
	// read stdin
	if len(files) == 0 {
		logger.Debugf("no files are given, reading stdin")
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			logger.Fatalf("failed to read stdin: %s", err)
		}
		b = bytes.TrimSpace(b)
		if len(b) > 0 {
			if split {
				s := bytes.Split(b, []byte(";"))
				for _, ss := range s {
					ss = bytes.TrimSpace(ss)
					if len(ss) > 0 {
						queries = append(queries, string(ss))
					}
				}
			} else {
				queries = append(queries, string(b))
			}
		}

		return
	}

	// read files
	for _, f := range files {
		logger.Debugf("reading file %s", f)
		b, err := ioutil.ReadFile(f) //nolint:gosec
		if err != nil {
			logger.Fatalf("failed to read file %s: %s", f, err)
		}
		b = bytes.TrimSpace(b)
		if len(b) > 0 {
			if split {
				s := bytes.Split(b, []byte(";"))
				for _, ss := range s {
					ss = bytes.TrimSpace(ss)
					if len(ss) > 0 {
						queries = append(queries, string(ss))
					}
				}
			} else {
				queries = append(queries, string(b))
			}
		}
	}

	return
}

// cmdExec implements exec command.
func cmdExec(db *reform.DB, files []string) {
	queries := readFiles(files, *execSplitF)
	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			logger.Fatalf("failed to execute %s: %s", q, err)
		}
	}
}
