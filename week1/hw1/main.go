package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func printNode(output io.Writer, info os.FileInfo, printFiles bool, isLast bool, prefix string) {
	jumper := "├"
	if isLast {
		jumper = "└"
	}
	branch := jumper + strings.Repeat("─", 3)
	fmt.Fprintf(output, "%s%s%s", prefix, branch, info.Name())
	if !info.IsDir() {
		formattedSize := fmt.Sprintf("%db", info.Size())
		if info.Size() == 0 {
			formattedSize = "empty"
		}
		fmt.Fprintf(output, " (%s)", formattedSize)
	}
	fmt.Fprintln(output)
}

func getChildren(file os.File, skipFiles bool) ([]os.FileInfo, error) {
	children, err := file.Readdir(0)
	if err != nil {
		return []os.FileInfo{}, err
	}

	if skipFiles {
		filtered := []os.FileInfo{}
		for _, child := range children {
			if child.IsDir() {
				filtered = append(filtered, child)
			}
		}
		children = filtered
	}
	sort.Slice(children, func(i, j int) bool {
		return children[i].Name() < children[j].Name()
	})
	return children, nil
}

func walk(output io.Writer, path string, printFiles bool, isLast bool, prefix string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	printNode(output, fileInfo, printFiles, isLast, prefix)
	if !fileInfo.IsDir() {
		return nil
	}

	children, err := getChildren(*file, !printFiles)
	if err != nil {
		return err
	}
	for pos, child := range children {
		childPath := filepath.Join(path, child.Name())
		isLastChild := pos == len(children)-1

		childPrefix := prefix + "│\t"
		if isLast {
			childPrefix = prefix + "\t"
		}

		err = walk(output, childPath, printFiles, isLastChild, childPrefix)
		if err != nil {
			return err
		}
	}

	return nil
}

func dirTree(output io.Writer, path string, printFiles bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	children, err := getChildren(*file, !printFiles)
	for pos, child := range children {
		childPath := filepath.Join(path, child.Name())
		isLast := pos == len(children)-1

		err = walk(output, childPath, printFiles, isLast, "")
		if err != nil {
			return err
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
