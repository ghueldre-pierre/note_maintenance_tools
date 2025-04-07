package dead_link_check

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "error : please provide root directory's path.")
		os.Exit(1)
	}
	///////////
	f, err := os.OpenFile("root_directory_list.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	///////////
	rootAbsolutePath, err := filepath.Abs(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error : %s\n", err.Error())
		os.Exit(1)
	}
	fsys := os.DirFS(rootAbsolutePath)
	var fileInfo fs.FileInfo
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fileInfo, err = d.Info()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(os.Stdout, "relative path : %s, name : %s, modfication date : %s\n", p, d.Name(), fileInfo.ModTime())
		return nil
	})
}
