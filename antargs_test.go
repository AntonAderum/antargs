package antargs

import (
	"testing"
)

func TestNewShouldInitialize(t *testing.T) {
	want := &AntArg{
		name: "test",
		help: "help_test",
		args: []*Arg{},
	}

	got, err := New("test", "help_test")

	if err != nil {
		t.Errorf("Got error from New: %s\n", err.Error())
	}

	if !got.Equal(*want) {
		t.Errorf(ExpectedGotAntArg(*want, *got))
	}
}

func TestNewShouldRejectNoName(t *testing.T) {
	_, err := New("", "help_test")

	if err == nil {
		t.Errorf(ExpectedGotString("error", "nil"))
	}
}

func TestNewArgShouldGiveNewArg(t *testing.T) {
	want := &AntArg{
		name: "test",
		help: "help_test",
		args: []*Arg{{help: "sub_help", name: "sub_name", isFlag: false, shortcut: "s", subArgs: []*Arg{}}},
	}

	got, _ := New("test", "help_test")

	got.NewArg("sub_name", "sub_help", false, "s")

	if !got.Equal(*want) {
		t.Errorf(ExpectedGotAntArg(*want, *got))
	}
}

func TestNewSubArgShouldGiveNewSubArg(t *testing.T) {
	want := &AntArg{
		name: "test",
		help: "help_test",
		args: []*Arg{
			{
				help:     "sub_help",
				name:     "sub_name",
				isFlag:   false,
				shortcut: "s",
				subArgs: []*Arg{{
					name:     "sub_sub_name",
					help:     "sub_sub_help",
					isFlag:   true,
					shortcut: "",
				}},
			},
		},
	}

	got, _ := New("test", "help_test")

	arg := got.NewArg("sub_name", "sub_help", false, "s")
	arg.NewSubArg("sub_sub_name", "sub_sub_help", true, "")

	if !got.Equal(*want) {
		t.Errorf(ExpectedGotAntArg(*want, *got))
	}
}
