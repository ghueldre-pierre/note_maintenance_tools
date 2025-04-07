package main

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"nmt": handleUserArgs,
	})
}

func TestInit(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Files: []string{
			"testscripts/init.txt",
		},
	})
}
