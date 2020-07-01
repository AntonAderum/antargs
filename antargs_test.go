package antargs

import (
	"fmt"
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

func TestNewArgShouldGiveNewArg(t *testing.T) {
	want := &AntArg{
		name: "test",
		help: "help_test",
		args: []*Arg{{help: "sub_help", name: "sub_name", isFlag: false, shortcut: "s", subArgs: []*Arg{}}},
	}

	got, _ := New("test", "help_test")

	got.NewArg("sub_name", "sub_help", false, "s")

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

	arg := got.NewArg("sub_name", "sub_help", false, "s")
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

	arg := got.NewArg("sub_name", "sub_help", false, "s")
	subArg := arg.NewSubArg("sub_sub_name", "sub_sub_help", true, "")
	subArg.NewSubArg("sub_sub_sub_name", "sub_sub_sub_help", false, "p")

	if !got.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *got))
	}
}

func TestPrintsArgumentInformationCorrectly(t *testing.T) {
	// We include more than we are testing to make sure we are only
	// getting what we want
	antArg, _ := New("test", "help_test")
	arg := antArg.NewArg("sub_name", "sub_help", false, "s")
	subArg := arg.NewSubArg("sub_sub_name", "sub_sub_help", true, "")
	subArg.NewSubArg("sub_sub_sub_name", "sub_sub_sub_help", false, "p")

	wanted := fmt.Sprintf("\n%s\n%s\n\nArguments:\n\n", "test", "help_test")
	wanted = wanted + fmt.Sprintf("\t%s:\t%s\n\t\t\t%s\n\t\t\tflag: %t\n", "sub_name", "sub_help", "s", false)
	got := getArgumentInformation(*antArg)
	if got != wanted {
		t.Errorf(expectedGotString(got, wanted))
	}
}

func TestPrintsSubArgumentInformationCorrectly(t *testing.T) {
	// We include more than we are testing to make sure we are only
	// getting what we want
	antArg, _ := New("test", "help_test")
	arg := antArg.NewArg("sub_name", "sub_help", false, "s")
	subArg := arg.NewSubArg("sub_sub_name", "sub_sub_help", true, "")
	subArg.NewSubArg("sub_sub_sub_name", "sub_sub_sub_help", false, "p")

	wanted := fmt.Sprintf(mainArgumentInformationFormat, "sub_name", "sub_help")
	wanted = wanted + fmt.Sprintf(subArgumentInformationFormat, "sub_sub_name", "sub_sub_help", "", true)
	got := getSubArgumentInformation(*antArg.args[0])
	if got != wanted {
		t.Errorf(expectedGotString(got, wanted))
	}
}

func TestPrintsSubSubArgumentInformationCorrectly(t *testing.T) {
	antArg, _ := New("test", "help_test")
	arg := antArg.NewArg("sub_name", "sub_help", false, "s")
	subArg := arg.NewSubArg("sub_sub_name", "sub_sub_help", true, "")
	subArg.NewSubArg("sub_sub_sub_name", "sub_sub_sub_help", false, "p")

	wanted := fmt.Sprintf(mainArgumentInformationFormat, "sub_sub_name", "sub_sub_help")
	wanted = wanted + fmt.Sprintf(subArgumentInformationFormat, "sub_sub_sub_name", "sub_sub_sub_help", "p", false)
	got := getSubArgumentInformation(*antArg.args[0].subArgs[0])
	if got != wanted {
		t.Errorf(expectedGotString(got, wanted))
	}
}
