package antargs

import (
	"testing"
)

func TestParseOneArgumentAndValue(t *testing.T) {
	want := getTestingAntArgObjectWithArgument(1)
	want.args[0].values = []string{"sub_value"}
	want.args[0].wasProvided = true

	antArg, _ := New("test", "help_test")
	antArg.NewArg("sub_name_0", "sub_help_0", false, "", 1)

	antArg.Parse([]string{"/test/test", "sub_name_0", "sub_value"})

	if !antArg.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *antArg))
	}
}

func TestParseOneArgumentAndMultipleValues(t *testing.T) {
	want := getTestingAntArgObjectWithArgument(1)
	want.args[0].values = []string{"sub_value_0", "sub_value_1"}
	want.args[0].numberOfValues = 2
	want.args[0].wasProvided = true

	antArg, _ := New("test", "help_test")
	antArg.NewArg("sub_name_0", "sub_help_0", false, "", 2)

	antArg.Parse([]string{"/test/test", "sub_name_0", "sub_value_0", "sub_value_1"})

	if !antArg.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *antArg))
	}
}

func TestParseOneArgumentAndNoValueProvidedReturnsError(t *testing.T) {
	want := getTestingAntArgObjectWithArgument(1)
	want.args[0].values = []string{"sub_value"}

	antArg, _ := New("test", "help_test")
	antArg.NewArg("sub_name_0", "sub_help_0", false, "", 1)

	err := antArg.Parse([]string{"/test/test", "sub_name_0"})
	if err == nil {
		t.Errorf(expectedGotString("err", "nil"))
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

func TestParseOneArgumentWithOneSubArgumentFlag(t *testing.T) {
	want := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	want.args[0].values = []string{"sub_value"}
	want.args[0].wasProvided = true
	want.args[0].subArgs[0].wasProvided = true
	want.args[0].subArgs[0].isFlag = true
	want.args[0].subArgs[0].numberOfValues = 0
	want.args[0].subArgs[0].values = []string{}

	antArg, _ := New("test", "help_test")
	subArg, _ := antArg.NewArg("sub_name_0", "sub_help_0", false, "", 1)
	subArg.NewSubArg("sub_sub_name_0", "sub_sub_help_0", true, "", 0)

	antArg.Parse([]string{"/test/test", "sub_name_0", "sub_sub_name_0", "sub_value"})

	if !antArg.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *antArg))
	}
}

func TestParseOneArgumentWithOneSubArgumentAndValue(t *testing.T) {
	want := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	want.args[0].values = []string{"sub_value"}
	want.args[0].wasProvided = true
	want.args[0].subArgs[0].wasProvided = true
	want.args[0].subArgs[0].numberOfValues = 1
	want.args[0].subArgs[0].values = []string{"sub_sub_value"}

	antArg, _ := New("test", "help_test")
	subArg, _ := antArg.NewArg("sub_name_0", "sub_help_0", false, "", 1)
	subArg.NewSubArg("sub_sub_name_0", "sub_sub_help_0", false, "", 1)

	antArg.Parse([]string{"/test/test", "sub_name_0", "sub_sub_name_0", "sub_sub_value", "sub_value"})

	if !antArg.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *antArg))
	}
}

func TestParseOneArgumentWithOneSubArgumentAndMultipleValues(t *testing.T) {
	want := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	want.args[0].values = []string{"sub_value"}
	want.args[0].wasProvided = true
	want.args[0].subArgs[0].wasProvided = true
	want.args[0].subArgs[0].numberOfValues = 2
	want.args[0].subArgs[0].values = []string{"sub_sub_value_0", "sub_sub_value_1"}

	antArg, _ := New("test", "help_test")
	subArg, _ := antArg.NewArg("sub_name_0", "sub_help_0", false, "", 1)
	subArg.NewSubArg("sub_sub_name_0", "sub_sub_help_0", false, "", 2)

	antArg.Parse([]string{"/test/test", "sub_name_0", "sub_sub_name_0", "sub_sub_value_0", "sub_sub_value_1", "sub_value"})

	if !antArg.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *antArg))
	}
}

func TestParseOneArgumentWithOneSubArgumentAndOnlyOneValueProvidedReturnsError(t *testing.T) {
	want := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	want.args[0].values = []string{"sub_value"}
	want.args[0].wasProvided = true
	want.args[0].subArgs[0].wasProvided = true
	want.args[0].subArgs[0].numberOfValues = 1
	want.args[0].subArgs[0].values = []string{"sub_sub_value"}

	antArg, _ := New("test", "help_test")
	subArg, _ := antArg.NewArg("sub_name_0", "sub_help_0", false, "", 1)
	subArg.NewSubArg("sub_sub_name_0", "sub_sub_help_0", false, "", 1)

	err := antArg.Parse([]string{"/test/test", "sub_name_0", "sub_sub_name_0", "sub_value"})

	if err == nil {
		t.Errorf(expectedGotString("err", "nil"))
	}
}

func TestParseOneArgumentWithOneSubArgumentAndNoValueProvidedReturnsError(t *testing.T) {
	want := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	want.args[0].values = []string{"sub_value"}
	want.args[0].wasProvided = true
	want.args[0].subArgs[0].wasProvided = true
	want.args[0].subArgs[0].numberOfValues = 1
	want.args[0].subArgs[0].values = []string{"sub_sub_value"}

	antArg, _ := New("test", "help_test")
	subArg, _ := antArg.NewArg("sub_name_0", "sub_help_0", false, "", 1)
	subArg.NewSubArg("sub_sub_name_0", "sub_sub_help_0", false, "", 1)

	err := antArg.Parse([]string{"/test/test", "sub_name_0", "sub_sub_name_0"})

	if err == nil {
		t.Errorf(expectedGotString("err", "nil"))
	}
}
