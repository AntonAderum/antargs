package antargs

import (
	"testing"
)

func TestCompareSameIsTrue(t *testing.T) {
	a := &AntArg{
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

	b := &AntArg{
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

	if !a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}

func TestCompareDifferentNameIsFalse(t *testing.T) {
	a := &AntArg{
		name: "test_different",
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

	b := &AntArg{
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

	if a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}

func TestCompareDifferentSubHelpIsFalse(t *testing.T) {
	a := &AntArg{
		name: "test",
		help: "help_test",
		args: []*Arg{
			{
				help:     "sub_help",
				name:     "sub_name",
				isFlag:   false,
				shortcut: "s",
				subArgs: []*Arg{{
					name:     "sub_sub_name_different",
					help:     "sub_sub_help",
					isFlag:   true,
					shortcut: "",
				}},
			},
		},
	}

	b := &AntArg{
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

	if a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}

func TestCompareDifferentSubLengthIsFalse(t *testing.T) {
	a := &AntArg{
		name: "test",
		help: "help_test",
		args: []*Arg{
			{
				help:     "sub_help",
				name:     "sub_name",
				isFlag:   false,
				shortcut: "s",
				subArgs:  []*Arg{},
			},
		},
	}

	b := &AntArg{
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

	if a.Equal(*b) {
		t.Errorf(expectedGotString("true", "false"))
	}
}
