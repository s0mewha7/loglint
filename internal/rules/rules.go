package rules

import (
	"go/ast"
	"regexp"
	"strings"
	"unicode"

	"github.com/s0mewha7/loglint/internal/utils"

	"golang.org/x/tools/go/analysis"
)

var specialChars = regexp.MustCompile(`[!@#$%^&*(){}\[\]<>🚀…]+`)

var sensitivePatterns = []string{
	"password",
	"passwd",
	"token",
	"api_key",
	"apikey",
	"secret",
}

func CheckLogContent(pass *analysis.Pass, call *ast.CallExpr) {
	if !utils.IsLogCall(call) {
		return
	}

	msg, ok := utils.GetMessage(call)
	if !ok {
		return
	}

	checkLowercase(pass, call, msg)
	checkEnglish(pass, call, msg)
	checkSpecialChars(pass, call, msg)
}

func CheckSecrets(pass *analysis.Pass, call *ast.CallExpr) {
	for _, arg := range call.Args {
		bin, ok := arg.(*ast.BinaryExpr)
		if !ok {
			continue
		}

		ident, ok := bin.Y.(*ast.Ident)
		if !ok {
			continue
		}

		name := strings.ToLower(ident.Name)

		for _, s := range sensitivePatterns {
			if strings.Contains(name, s) {
				pass.Reportf(call.Pos(),
					"possible sensitive data in logs: %s", ident.Name)
			}
		}
	}
}

func checkLowercase(pass *analysis.Pass, node ast.Node, msg string) {
	if len(msg) == 0 {
		return
	}

	r := rune(msg[0])

	if unicode.IsUpper(r) {
		pass.Reportf(node.Pos(),
			"log message must start with lowercase")
	}
}

func checkEnglish(pass *analysis.Pass, node ast.Node, msg string) {
	for _, r := range msg {
		if unicode.IsLetter(r) && r > unicode.MaxASCII {
			pass.Reportf(node.Pos(),
				"log message must be in english")
			return
		}
	}
}

func checkSpecialChars(pass *analysis.Pass, node ast.Node, msg string) {
	if specialChars.MatchString(msg) {
		pass.Reportf(node.Pos(),
			"log message must not contain special characters")
	}
}
