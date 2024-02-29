package main

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/qw4n7y/reform"
	"github.com/qw4n7y/reform/internal"
	"github.com/qw4n7y/reform/internal/test"
)

type ReformDBSuite struct {
	suite.Suite
	db *reform.DB
}

func TestReformDBSuite(t *testing.T) {
	suite.Run(t, new(ReformDBSuite))
}

func (s *ReformDBSuite) SetupSuite() {
	if testing.Short() {
		s.T().Skip("skipping in short mode")
	}

	logger = internal.NewLogger("reform-db-test: ", true)

	s.db = test.ConnectToTestDB()
	s.db.Querier = s.db.WithTag("reform-db-test")
}

func (s *ReformDBSuite) SetupTest() {
	pl := reform.NewPrintfLogger(s.T().Logf)
	pl.LogTypes = true
	s.db.Logger = pl
}
