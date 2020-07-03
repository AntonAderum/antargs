package antargs

import (
	"encoding/json"
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func deepToJSON(args []*Arg) []jsonArg {
	jsonArgs := make([]jsonArg, len(args))
	for i, arg := range args {
		jsonArgs[i] = jsonArg{
			Name:           arg.name,
			Help:           arg.help,
			Shortcut:       arg.shortcut,
			IsFlag:         arg.isFlag,
			SubArgs:        deepToJSON(arg.subArgs),
			Values:         arg.values,
			NumberOfValues: arg.numberOfValues,
			WasProvided:    arg.wasProvided,
		}
	}
	return jsonArgs
}

// Prettify returns an JSON object as a string representing the AntArg
func (antArg AntArg) Prettify() string {
	i := jsonAntArg{
		Name: antArg.name,
		Help: antArg.help,
		Args: deepToJSON(antArg.args),
	}
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}

var (
	mainArgumentInformationFormat = "\n%s\n%s\n\nArguments:\n\n"
	subArgumentInformationFormat  = "\t%s:\t%s\n\t\t\tshortcut: %s\n\t\t\tflag: %t\n"
)

func getSubArgumentInformation(arg Arg) string {
	info := fmt.Sprintf(mainArgumentInformationFormat, arg.name, arg.help)
	for _, arg := range arg.subArgs {
		info = info + fmt.Sprintf(subArgumentInformationFormat, arg.name, arg.help, arg.shortcut, arg.isFlag)
	}
	return info
}

func getArgumentInformation(antArg AntArg) string {
	info := fmt.Sprintf("\n%s\n%s\n\nArguments:\n\n", antArg.name, antArg.help)
	for _, arg := range antArg.args {
		info = info + fmt.Sprintf("\t%s:\t%s\n\t\t\t%s\n\t\t\tflag: %t\n", arg.name, arg.help, arg.shortcut, arg.isFlag)
	}
	return info
}

func (arg Arg) PrintSubArgumentInformation() {
	fmt.Print(getSubArgumentInformation(arg))
}

func (antArg AntArg) PrintArgumentInformation() {
	fmt.Print(getArgumentInformation(antArg))
}

// expectedGotAntArg returns a formatted string for
// the normal use case of "expected a got b" string.
// It prints the AntArg objects as JSON objecs, and also
// diff and prints the diff for easy identification of the problem
func expectedGotAntArg(a AntArg, b AntArg) string {
	prettyA := a.Prettify()
	prettyB := b.Prettify()
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(prettyA, prettyB, true)
	return fmt.Sprintf("\nexpected:\n%s\ngot:\n%s\ndiff:\n%s\n", prettyA, prettyB, dmp.DiffPrettyText(diffs))
}

// ExpectedGotString returns a formatted string for
// the normal use case of "expected a got b" string
func expectedGotString(a string, b string) string {
	return fmt.Sprintf("\nexpected: \"%s\" got: \"%s\"\n", a, b)
}
