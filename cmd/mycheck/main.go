package main

import (
	"honnef.co/go/tools/staticcheck"

	"github.com/salihzain/tagalyzer"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
)

func main() {
	var mychecks []*analysis.Analyzer
	mychecks = append(
		mychecks,
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		tagalyzer.Analyzer,
	)

	checks := map[string]bool{
		"SA":     true,
		"ST1008": true,
	}
	for _, v := range staticcheck.Analyzers {
		if checks[v.Analyzer.Name] {
			mychecks = append(mychecks, v.Analyzer)
		}
	}

	multichecker.Main(
		mychecks...,
	)
}
