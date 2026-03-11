package analyzer

import (
	"go/ast"

	"github.com/s0mewha7/loglint/internal/rules"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "loglint",
		Doc:      "checks log messages style and secrets",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {

		call := n.(*ast.CallExpr)

		rules.CheckLogContent(pass, call)
		rules.CheckSecrets(pass, call)

	})

	return nil, nil
}
