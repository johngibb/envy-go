package envy

import (
	"os"
	"reflect"
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
	var exitCalled bool
	exit = func(int) {
		exitCalled = true
	}
	MustLoad(".env.notfound")
	assertEq(t, "exitCalled", exitCalled, true)
}

func assertEq(t *testing.T, field string, got, want any) {
	if !reflect.DeepEqual(got, want) {
		t.Helper()
		t.Errorf("%s: got \"%v\", want \"%v\"", field, got, want)
	}
}
