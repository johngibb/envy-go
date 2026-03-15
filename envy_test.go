package envy

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	assertEq(t,
		`os.Getenv("ENVY_APP_NAME")`,
		os.Getenv("ENVY_APP_NAME"),
		"",
	)
	Load(".env.example")
	assertEq(t,
		`os.Getenv("ENVY_APP_NAME")`,
		os.Getenv("ENVY_APP_NAME"),
		"envy",
	)
}
func TestMustLoad(t *testing.T) {
	var logFatalfCalled bool
	logFatalf = func(format string, v ...any) {
		logFatalfCalled = true
	}
	badFileName := strings.Repeat("too:long", 40)
	MustLoad(badFileName)
	assertEq(t, "logFatalfCalled", logFatalfCalled, true)
}

func assertEq(t *testing.T, field string, got, want any) {
	if !reflect.DeepEqual(got, want) {
		t.Helper()
		t.Errorf("%s: got \"%v\", want \"%v\"", field, got, want)
	}
}
