package main

import (
	"github.com/artemgareev/logkeyslint/pkg/linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(linter.Analyzer)
}
