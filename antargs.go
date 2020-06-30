package antargs

import (
	"fmt"
)

type jsonArg struct {
	Name     string
	Help     string
	Shortcut string
	IsFlag   bool
	SubArgs  []jsonArg
}

type jsonAntArg struct {
	Name string
	Help string
	Args []jsonArg
}

// Arg is the struct containing a specific argument for an AntArg list of arguments
type Arg struct {
	name     string
	help     string
	shortcut string
	isFlag   bool
	subArgs  []*Arg
}

// AntArg is the main struct when working with AntArg package.
// It contains an array of arguments of the using program,
// as well as information about the program.
type AntArg struct {
	name string
	help string
	args []*Arg
}

// New initializes an instance of AntArgs, should be used for creating the main program AntArg instance
func New(name string, help string) (*AntArg, error) {

	if len(name) == 0 {
		return nil, fmt.Errorf("Name must have a value")
	}

	return &AntArg{
		name: name,
		help: help,
		args: []*Arg{},
	}, nil
}

// NewSubArg initializes a new argument tied to a parent arg
func (arg *Arg) NewSubArg(name string, help string, isFlag bool, shortcut string) {
	subArg := &Arg{
		name:     name,
		help:     help,
		isFlag:   isFlag,
		shortcut: shortcut,
	}
	arg.subArgs = append(arg.subArgs, subArg)
}

// NewArg initializes a new argument tied to a parent AntArg
func (antArg *AntArg) NewArg(name string, help string, isFlag bool, shortcut string) *Arg {
	arg := &Arg{
		name:     name,
		help:     help,
		isFlag:   isFlag,
		shortcut: shortcut,
	}
	antArg.args = append(antArg.args, arg)
	return arg
}
