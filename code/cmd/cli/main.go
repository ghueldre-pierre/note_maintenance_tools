package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const NMT_DIR = ".nmt"

func main() {
	handleUserArgs()
}

func handleUserArgs() {
	if len(os.Args) < 2 {
		fmt.Println("TODO : show help when no arguments are given")
		os.Exit(0)
	}
	{
		cmd := os.Args[1]
		if cmd == "init" {
			err := handleInit()
			if err != nil {
				fmt.Fprintln(os.Stderr, "error : ", err.Error())
				os.Exit(1)
			}
		}
	}
}

func handleInit() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	nmtDirPath := filepath.Join(wd, NMT_DIR)
	_, err = os.Stat(nmtDirPath)
	if errors.Is(err, fs.ErrNotExist) {
		err = os.Mkdir(nmtDirPath, 0750)
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, "Note Maintenance Tools has been successfully initialized.")
	} else {
		fmt.Fprintln(os.Stdout, "Note Maintenance Tools has already been initialized.")
	}
	return nil
}
