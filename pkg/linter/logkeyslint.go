package linter

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var logTypes = []string{
	"Str", "Strs", "Bytes", "Hex", "Errs", "Err", "Bool", "Bools",
	"Int", "Ints", "Int8", "Ints8", "Int16", "Ints16", "Int32", "Ints32", "Int64", "Ints64",
	"Uint", "Uints", "Uint8", "Uints8", "Uint16", "Uints16", "Uint32", "Uints32", "Uint64", "Uints64",
	"Time", "Times", "Dur", "Durs",
}

var intTypes = []string{
	"Int", "Int8", "Int16", "Int32", "Int64", "Uint", "Uint8", "Uint16", "Uint32", "Uint64",
}

var intsTypes = []string{
	"Ints", "Ints8", "Ints16", "Ints32", "Ints64", "Uints", "Uints8", "Uints16", "Uints32", "Uints64",
}

var logFieldMap = map[string]string{}

var Analyzer = &analysis.Analyzer{
	Name: "logkeyslint",
	Doc:  "Checks rs/zerolog log keys types mismatches",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		funcDecl, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		fun, ok := funcDecl.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		if !inArrayString(fun.Sel.Name, logTypes) {
			return true
		}

		if len(funcDecl.Args) < 1 {
			return true
		}
		arg, ok := funcDecl.Args[0].(*ast.BasicLit)
		if !ok {
			return true
		}

		if prevType, ok := logFieldMap[arg.Value]; ok {
			if prevType != fun.Sel.Name {
				if inArrayString(prevType, intTypes) && inArrayString(fun.Sel.Name, intTypes) {
					return true
				}
				if inArrayString(prevType, intsTypes) && inArrayString(fun.Sel.Name, intsTypes) {
					return true
				}
				pass.Reportf(node.Pos(), "Bad %s log key type of \"%s\" previously was used as \"%s\"\n",
					arg.Value, fun.Sel.Name, prevType)
				return true
			}
		} else {
			logFieldMap[arg.Value] = fun.Sel.Name
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}

func inArrayString(value string, haystack []string) bool {
	for _, v := range haystack {
		if value == v {
			return true
		}
	}
	return false
}
