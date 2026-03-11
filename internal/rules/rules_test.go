package rules

import "testing"

func TestLowercase(t *testing.T) {
	cases := map[string]bool{
		"Starting server": true,
		"starting server": false,
		"Failed":          true,
		"failed":          false,
		"":                false,
	}

	for msg, want := range cases {
		if got := isUppercase(msg); got != want {
			t.Errorf("isUppercase(%q) = %v, want %v", msg, got, want)
		}
	}
}

func TestEnglish(t *testing.T) {
	cases := map[string]bool{
		"запуск сервера":  true,
		"starting server": false,
		"ошибка":          true,
		"error occurred":  false,
	}

	for msg, want := range cases {
		if got := isNonEnglish(msg); got != want {
			t.Errorf("isNonEnglish(%q) = %v, want %v", msg, got, want)
		}
	}
}

func TestSpecialChars(t *testing.T) {
	cases := map[string]bool{
		"server started!":      true,
		"failed!!!":            true,
		"server started🚀":      true,
		"something went wrong": false,
		"server started":       false,
	}

	for msg, want := range cases {
		if got := hasSpecialChars(msg); got != want {
			t.Errorf("hasSpecialChars(%q) = %v, want %v", msg, got, want)
		}
	}
}

func TestSecrets(t *testing.T) {
	cases := map[string]bool{
		"password":    true,
		"apiKey":      true,
		"token":       true,
		"userID":      false,
		"requestBody": false,
	}

	for s, want := range cases {
		if got := isSensitive(s); got != want {
			t.Errorf("isSensitive(%q) = %v, want %v", s, got, want)
		}
	}
}
