package antargs

import (
	"testing"
)

func TestCompareSameIsTrue(t *testing.T) {
	a := getTestingAntArgObjectWithArgumentAndSubArgumentsAndSubArguments(1, 1, 1)

	b := getTestingAntArgObjectWithArgumentAndSubArgumentsAndSubArguments(1, 1, 1)

	if !a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}

func TestCompareDifferentNameIsFalse(t *testing.T) {
	a := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	a.name = "test_different"

	b := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)

	if a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}

func TestCompareDifferentSubArgumentShortcutIsFalse(t *testing.T) {
	a := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	a.args[0].subArgs[0].shortcut = "d"

	b := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)

	if a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}

func TestCompareDifferentSubArgumentNameIsFalse(t *testing.T) {
	a := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	a.args[0].subArgs[0].name = "different"

	b := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)

	if a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}

func TestCompareDifferentSubArgumentHelpIsFalse(t *testing.T) {
	a := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	a.args[0].subArgs[0].help = "different"

	b := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)

	if a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}

func TestCompareDifferentSubArgumentIsFlagIsFalse(t *testing.T) {
	a := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	a.args[0].subArgs[0].isFlag = true

	b := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)

	if a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}

func TestCompareDifferentSubArgumentNumberOfValuesIsFalse(t *testing.T) {
	a := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	a.args[0].subArgs[0].numberOfValues = 2

	b := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)

	if a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}

func TestCompareDifferentSubArgumentWasProvidedIsFalse(t *testing.T) {
	a := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)
	a.args[0].subArgs[0].wasProvided = true

	b := getTestingAntArgObjectWithArgumentAndSubArguments(1, 1)

	if a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}

func TestCompareDifferentSubLengthIsFalse(t *testing.T) {
	a := getTestingAntArgObjectWithArgumentAndSubArgumentsAndSubArguments(1, 1, 2)

	b := getTestingAntArgObjectWithArgumentAndSubArgumentsAndSubArguments(1, 1, 1)

	if a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}
