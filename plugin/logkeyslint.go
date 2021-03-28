package main

import (
	"github.com/artemgareev/logkeyslint/pkg/linter"
	"golang.org/x/tools/go/analysis"
)

// ref: https://golangci-lint.run/contributing/new-linters/#how-to-add-a-private-linter-to-golangci-lint
type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		linter.Analyzer,
	}
}

var AnalyzerPlugin analyzerPlugin
