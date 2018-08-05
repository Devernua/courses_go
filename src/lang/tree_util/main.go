package main

import (
	"io"
	"os"
	"fmt"
	"sort"
	"bytes"
)

func dirTree(output io.Writer, path string, printFiles bool) error {
	dir, err := os.Open(path)
	if err != nil {
		return err
	}

	infos, err := dir.Readdir(0)
	if err != nil {
		return err
	}

	// filtering files if need
	if !printFiles {
		var newInfos []os.FileInfo
		for _, file := range infos {
			if file.IsDir() {
				newInfos = append(newInfos, file)
			}
		}
		infos = newInfos
	}

	// sorting by lex
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Name() < infos[j].Name()
	})

	// print
	for i, file := range infos {
		sep := "├───"
		sep2 := "│\t"
		if i + 1 == len(infos) {
			sep = "└───"
			sep2 = "\t"
		}

		if file.IsDir() {
			fmt.Fprintln(output, sep + file.Name())
			out := new(bytes.Buffer)
			err = dirTree(out, path + "/" + file.Name(), printFiles)
			if err != nil {
				return err
			}

			row, err := out.ReadString('\n')

			//rowSize, row, err := bufio.ScanLines(out.Bytes(), true)
			for len(row) != 0 && err == nil {
				fmt.Fprint(output, sep2 + row)
				row, err = out.ReadString('\n')
			}
		} else {
			strSize := "("
			if file.Size() == 0 {
				strSize += "empty)"
			} else {
				strSize += fmt.Sprintf("%vb)", file.Size())
			}

			fmt.Fprintln(output, sep + file.Name() + " " + strSize)
		}
	}

	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
