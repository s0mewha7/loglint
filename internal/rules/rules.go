package rules

import (
	"go/ast"
	"regexp"
	"strings"
	"unicode"

	"github.com/s0mewha7/loglint/internal/utils"
	"golang.org/x/tools/go/analysis"
)

var badChars = regexp.MustCompile(`[!@#$%^&*(){}\[\]<>🚀…]+`)

var secrets = []string{
	"password", "passwd",
	"token", "secret",
	"api_key", "apikey",
}

func CheckLogContent(pass *analysis.Pass, call *ast.CallExpr) {
	if !utils.IsLogCall(call) {
		return
	}

	msg, ok := utils.GetMessage(call)
	if !ok {
		return
	}

	if len(msg) > 0 && unicode.IsUpper(rune(msg[0])) {
		pass.Reportf(call.Pos(), "log message must start with lowercase")
	}

	for _, r := range msg {
		if unicode.IsLetter(r) && r > unicode.MaxASCII {
			pass.Reportf(call.Pos(), "log message must be in english")
			break
		}
	}

	if badChars.MatchString(msg) {
		pass.Reportf(call.Pos(), "log message must not contain special characters")
	}
}

func CheckSecrets(pass *analysis.Pass, call *ast.CallExpr) {
	if !utils.IsLogCall(call) {
		return
	}

	for _, arg := range call.Args {
		checkExpr(pass, call, arg)
	}
}

func checkExpr(pass *analysis.Pass, call *ast.CallExpr, expr ast.Expr) {
	switch e := expr.(type) {
	case *ast.BinaryExpr:
		checkExpr(pass, call, e.X)
		checkExpr(pass, call, e.Y)
	case *ast.Ident:
		if isSensitive(e.Name) {
			pass.Reportf(call.Pos(), "possible sensitive data in logs: %s", e.Name)
		}
	case *ast.BasicLit:
		val := strings.Trim(e.Value, `"`)
		if isSensitive(val) {
			pass.Reportf(call.Pos(), "possible sensitive data in logs: %s", val)
		}
	}
}

func isSensitive(s string) bool {
	s = strings.ToLower(s)
	for _, secret := range secrets {
		if strings.Contains(s, secret) {
			return true
		}
	}
	return false
}

func isUppercase(msg string) bool {
	if len(msg) == 0 {
		return false
	}
	return unicode.IsUpper(rune(msg[0]))
}

func isNonEnglish(msg string) bool {
	for _, r := range msg {
		if unicode.IsLetter(r) && r > unicode.MaxASCII {
			return true
		}
	}
	return false
}

func hasSpecialChars(msg string) bool {
	return badChars.MatchString(msg)
}
