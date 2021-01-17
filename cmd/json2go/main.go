package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/ttyfky/json2go/descriptor"
	"github.com/ttyfky/json2go/internal/gengo"
)

const defaultErrorMsg = "json2go: %s\n"

var (
	outputPath = flag.String("output", "", "path to directory to put output. if not specified results are printed in stdout.")
	pkgName    = flag.String("pkg", "gen", "package name of generated Go file.")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Printf("usage: json2go [opts] PATH_TO_FIlE_OR_DIR\n" +
			"opts:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	reg, err := initRegistry(args[0])
	if err != nil {
		fmt.Printf(defaultErrorMsg, err)
		os.Exit(1)
	}
	files, err := reg.Load()
	if err != nil {
		fmt.Printf(defaultErrorMsg, err)
		os.Exit(1)
	}
	g := gengo.New(reg)
	err = g.Generate(files)
	if err != nil {
		fmt.Printf(defaultErrorMsg, err)
		os.Exit(1)
	}
}

func initRegistry(argPath string) (*descriptor.Registry, error) {
	var reg *descriptor.Registry
	{
		reg = descriptor.NewRegistry(*pkgName)
		err := reg.SetupInput(path.Clean(argPath))
		if err != nil {
			return nil, err
		}
	}
	if *outputPath != "" {
		if err := os.MkdirAll(*outputPath, 0755); err != nil {
			return nil, err
		}
		reg.SetOutputPath(*outputPath)
	} else {
		reg.SetWriter(os.Stdout)
	}

	return reg, nil
}
