package antargs

import (
	"fmt"
	"testing"
)

func TestPrintsArgumentInformationCorrectly(t *testing.T) {
	// We include more than we are testing to make sure we are only
	// getting what we want
	antArg, _ := New("test", "help_test")
	arg, _ := antArg.NewArg("sub_name", "sub_help", false, "s", 1)
	subArg, _ := arg.NewSubArg("sub_sub_name", "sub_sub_help", true, "", 1)
	subArg.NewSubArg("sub_sub_sub_name", "sub_sub_sub_help", false, "p", 1)

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
	arg, _ := antArg.NewArg("sub_name", "sub_help", false, "s", 1)
	subArg, _ := arg.NewSubArg("sub_sub_name", "sub_sub_help", true, "", 1)
	subArg.NewSubArg("sub_sub_sub_name", "sub_sub_sub_help", false, "p", 1)

	wanted := fmt.Sprintf(mainArgumentInformationFormat, "sub_name", "sub_help")
	wanted = wanted + fmt.Sprintf(subArgumentInformationFormat, "sub_sub_name", "sub_sub_help", "", true)
	got := getSubArgumentInformation(*antArg.args[0])
	if got != wanted {
		t.Errorf(expectedGotString(got, wanted))
	}
}

func TestPrintsSubSubArgumentInformationCorrectly(t *testing.T) {
	antArg, _ := New("test", "help_test")
	arg, _ := antArg.NewArg("sub_name", "sub_help", false, "s", 1)
	subArg, _ := arg.NewSubArg("sub_sub_name", "sub_sub_help", true, "", 1)
	subArg.NewSubArg("sub_sub_sub_name", "sub_sub_sub_help", false, "p", 1)

	wanted := fmt.Sprintf(mainArgumentInformationFormat, "sub_sub_name", "sub_sub_help")
	wanted = wanted + fmt.Sprintf(subArgumentInformationFormat, "sub_sub_sub_name", "sub_sub_sub_help", "p", false)
	got := getSubArgumentInformation(*antArg.args[0].subArgs[0])
	if got != wanted {
		t.Errorf(expectedGotString(got, wanted))
	}
}
