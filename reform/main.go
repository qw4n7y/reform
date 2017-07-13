// Package reform implements reform command.
package main

import (
	"flag"
	"fmt"
	"go/build"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/internal"
	"gopkg.in/reform.v1/parse"
)

var (
	logger *internal.Logger

	debugF = flag.Bool("debug", false, "Enable debug logging")
	gofmtF = flag.Bool("gofmt", true, "Format with gofmt")
)

func processFile(path, file, pack string) error {
	logger.Debugf("processFile: path=%q file=%q pack=%q", path, file, pack)

	structs, err := parse.File(filepath.Join(path, file))
	if err != nil {
		return err
	}

	logger.Debugf("%#v", structs)
	if len(structs) == 0 {
		return nil
	}

	ext := filepath.Ext(file)
	base := strings.TrimSuffix(file, ext)
	f, err := os.Create(filepath.Join(path, base+"_reform"+ext))
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString("// Code generated by gopkg.in/reform.v1. DO NOT EDIT.\n\n"); err != nil {
		return err
	}
	if _, err = f.WriteString("package " + pack + "\n"); err != nil {
		return err
	}
	if err = prologTemplate.Execute(f, nil); err != nil {
		return err
	}

	sds := make([]StructData, 0, len(structs))
	for _, str := range structs {
		// decide about view/table suffix
		t := strings.ToLower(str.Type[0:1]) + str.Type[1:]
		v := str.Type
		if str.IsTable() {
			t += "TableType"
			v += "Table"
		} else {
			t += "ViewType"
			v += "View"
		}

		sd := StructData{
			StructInfo: str,
			TableType:  t,
			TableVar:   v,
		}
		sds = append(sds, sd)

		if err = structTemplate.Execute(f, &sd); err != nil {
			return err
		}
	}

	return initTemplate.Execute(f, sds)
}

func gofmt(path string) {
	if *gofmtF {
		cmd := exec.Command("gofmt", "-s", "-w", path)
		logger.Debugf(strings.Join(cmd.Args, " "))
		b, err := cmd.CombinedOutput()
		if err != nil {
			logger.Fatalf("gofmt error: %s", err)
		}
		logger.Debugf("gofmt output: %s", b)
	}
}

func goformat(filePath string) {
	if !*gofmtF {
		in, err := ioutil.ReadFile(filePath)
		if err != nil {
			logger.Fatalf("go/format read error: %s", err)
		}
		out, err := format.Source(in)
		if err != nil {
			logger.Fatalf("go/format formatting error: %s", err)
		}
		if !reflect.DeepEqual(in, out) {
			if err := ioutil.WriteFile(filePath, out, 0644); err != nil {
				logger.Fatalf("go/format write error: %s", err)
			}
		}
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "reform - a better ORM generator. %s.\n\n", reform.Version)
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "  %s [flags] [packages or directories]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  go generate [flags] [packages or files] (with '//go:generate reform' in files)\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	logger = internal.NewLogger("reform: ", *debugF)

	logger.Debugf("Environment:")
	for _, pair := range os.Environ() {
		if strings.HasPrefix(pair, "GO") {
			logger.Debugf("\t%s", pair)
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		logger.Fatalf("%s", err)
	}
	logger.Debugf("wd: %s", wd)
	logger.Debugf("args: %v", flag.Args())

	// process arguments
	for _, arg := range flag.Args() {
		// import arg as directory or package path
		var pack *build.Package
		s, err := os.Stat(arg)
		if err == nil && s.IsDir() {
			pack, err = build.ImportDir(arg, 0)
		}
		if os.IsNotExist(err) {
			err = nil
		}
		if pack == nil && err == nil {
			pack, err = build.Import(arg, wd, 0)
		}
		if err != nil {
			logger.Fatalf("%s: %s", arg, err)
		}

		logger.Debugf("%#v", pack)

		var changed bool
		for _, f := range pack.GoFiles {
			err = processFile(pack.Dir, f, pack.Name)
			if err != nil {
				logger.Fatalf("%s %s: %s", arg, f, err)
			}
			goformat(filepath.Join(pack.Dir, f))
			changed = true
		}

		if changed {
			gofmt(pack.Dir)
		}
	}

	// process go generate environment
	file := os.Getenv("GOFILE")
	pack := os.Getenv("GOPACKAGE")
	if file != "" && pack != "" {
		err := processFile(wd, file, pack)
		if err != nil {
			logger.Fatalf("%s", err)
		}
		goformat(file)
		gofmt(wd)
	}
}
