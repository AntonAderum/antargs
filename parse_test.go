package antargs

import (
	"testing"
)

func TestParseOneArgumentAndValue(t *testing.T) {
	want := getTestingAntArgObjectWithArgument(1)
	want.args[0].values = []string{"sub_value"}

	antArg, _ := New("test", "help_test")
	antArg.NewArg("sub_name_0", "sub_help_0", false, "", 1)

	antArg.Parse([]string{"/test/test", "sub_name_0", "sub_value"})

	if !antArg.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *antArg))
	}
}

func TestParseRequireOneTopLeveArguementReturnsErrorWithNoArgument(t *testing.T) {
	antArg, _ := New("test", "help_test")
	antArg.NewArg("sub_name", "sub_help", false, "s", 1)

	err := antArg.Parse([]string{"/test/test"}, RequireAtleastOneTopLevelArgument())

	if err == nil {
		t.Errorf(expectedGotString("error", "nil"))
	}
}

func TestParseRequireOneTopLeveArguementReturnsNoErrorWithArgument(t *testing.T) {
	antArg, _ := New("test", "help_test")
	antArg.NewArg("sub_name", "sub_help", false, "s", 1)

	err := antArg.Parse([]string{"/test/test", "sub_name", "sub_value"}, RequireAtleastOneTopLevelArgument())

	if err != nil {
		t.Errorf(expectedGotString("nil", "error"))
	}
}

func TestParseAllowOnlyOneTopLevelReturnsErrorWithTwoArgument(t *testing.T) {
	antArg, _ := New("test", "help_test")
	antArg.NewArg("sub_name_0", "sub_help_0", false, "", 1)
	antArg.NewArg("sub_name_1", "sub_help_1", false, "", 1)

	err := antArg.Parse([]string{"/test/test", "sub_name_0", "sub_value_0", "sub_name_1", "sub_value_1"}, AllowOnlyOneTopLevelArgument())

	if err == nil {
		t.Errorf(expectedGotString("error", "nil"))
	}
}

func TestParseAllowOnlyOneTopLevelReturnsNoErrorWithOneArgument(t *testing.T) {
	antArg, _ := New("test", "help_test")
	antArg.NewArg("sub_name_0", "sub_help_0", false, "", 1)
	antArg.NewArg("sub_name_1", "sub_help_1", false, "", 1)

	err := antArg.Parse([]string{"/test/test", "sub_name_0", "sub_value"}, AllowOnlyOneTopLevelArgument())

	if err != nil {
		t.Errorf(expectedGotString("nil", "error"))
	}
}
