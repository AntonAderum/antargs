package antargs

import (
	"testing"
)

func TestNewShouldInitialize(t *testing.T) {
	want := getTestingAntArgObject()

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
	_, err := arg.NewSubArg("", "sub_sub_help", false, "", 1)

	if err == nil {
		t.Errorf(expectedGotString("error", "nil"))
	}
}

func TestNewArgShouldGiveNewArg(t *testing.T) {
	want := getTestingAntArgObjectWithArgument(1)

	got, _ := New("test", "help_test")

	got.NewArg("sub_name_0", "sub_help_0", false, "", 1)

	if !got.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *got))
	}
}

func TestNewSubArgShouldGiveNewSubArg(t *testing.T) {
	want := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)

	got, _ := New("test", "help_test")

	arg, _ := got.NewArg("sub_name_0", "sub_help_0", false, "", 1)
	arg.NewSubArg("sub_sub_name_0", "sub_sub_help_0", false, "", 1)

	if !got.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *got))
	}
}

func TestCanDoNestedSubArg(t *testing.T) {
	want := getTestingAntArgObjectWithArgumentAndSubArgumentsAndSubArguments(1, 1, 1)

	got, _ := New("test", "help_test")
	arg, _ := got.NewArg("sub_name_0", "sub_help_0", false, "", 1)
	subArg, _ := arg.NewSubArg("sub_sub_name_0", "sub_sub_help_0", false, "", 1)
	subArg.NewSubArg("sub_sub_sub_name_0", "sub_sub_sub_help_0", false, "", 1)

	if !got.Equal(*want) {
		t.Errorf(expectedGotAntArg(*want, *got))
	}
}
