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
		t.Errorf(expectedGotAntArg(*want, *got))
	}
}

func TestNewShouldRejectNoName(t *testing.T) {
	_, err := New("", "help_test")

	if err == nil {
		t.Errorf(expectedGotString("error", "nil"))
	}
}

func TestNewArgShouldRejectNoName(t *testing.T) {
	antArg, _ := New("test", "help_test")
	_, err := antArg.NewArg("", "sub_help", false, "", 1)

	if err == nil {
		t.Errorf(expectedGotString("error", "nil"))
	}
}

func TestNewSubArgShouldRejectNoName(t *testing.T) {
	antArg, _ := New("test", "help_test")
	arg, _ := antArg.NewArg("sub_test", "sub_help", false, "", 1)
	_, err := arg.NewSubArg("", "sub_sub_help", false, "")

	if err == nil {
		t.Errorf(expectedGotString("error", "nil"))
	}
}

func TestNewArgShouldGiveNewArg(t *testing.T) {
	want := &AntArg{
		name: "test",
		help: "help_test",
		args: []*Arg{{help: "sub_help", name: "sub_name", isFlag: false, shortcut: "s", subArgs: []*Arg{}}},
	}

	got, _ := New("test", "help_test")

	got.NewArg("sub_name", "sub_help", false, "s", 1)

	if !got.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *got))
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

	arg, _ := got.NewArg("sub_name", "sub_help", false, "s", 1)
	arg.NewSubArg("sub_sub_name", "sub_sub_help", true, "")

	if !got.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *got))
	}
}

func TestCanDoNestedSubArg(t *testing.T) {
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
					subArgs: []*Arg{{
						name:     "sub_sub_sub_name",
						help:     "sub_sub_sub_help",
						isFlag:   false,
						shortcut: "p",
					}},
				}},
			},
		},
	}

	got, _ := New("test", "help_test")

	arg, _ := got.NewArg("sub_name", "sub_help", false, "s", 1)
	subArg, _ := arg.NewSubArg("sub_sub_name", "sub_sub_help", true, "")
	subArg.NewSubArg("sub_sub_sub_name", "sub_sub_sub_help", false, "p")

	if !got.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *got))
	}
}
