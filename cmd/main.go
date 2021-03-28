package main

import (
	"github.com/artemgareev/zlogkeyscheck/pkg/linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(linter.Analyzer)
}
