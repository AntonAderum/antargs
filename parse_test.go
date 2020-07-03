package antargs

import (
	"testing"
)

func TestParseOneArgumentAndValue(t *testing.T) {
	want := &AntArg{
		name: "test",
		help: "help_test",
		args: []*Arg{{help: "sub_help", name: "sub_name", isFlag: false, shortcut: "s", numberOfValues: 1, values: []string{"sub_value"}, subArgs: []*Arg{}}},
	}

	antArg, _ := New("test", "help_test")
	antArg.NewArg("sub_name", "sub_help", false, "s", 1)

	antArg.Parse([]string{"/test/test", "sub_name", "sub_value"})

	if !antArg.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *antArg))
	}
}
