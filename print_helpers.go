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
			Name:     arg.name,
			Help:     arg.help,
			Shortcut: arg.shortcut,
			IsFlag:   arg.isFlag,
			SubArgs:  deepToJSON(arg.subArgs),
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
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// ExpectedGotAntArg returns a formatted string for
// the normal use case of "expected a got b" string
func ExpectedGotAntArg(a AntArg, b AntArg) string {
	prettyA := a.Prettify()
	prettyB := b.Prettify()
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(prettyA, prettyB, true)
	return fmt.Sprintf("\nexpected:\n%s\ngot:\n%s\ndiff:\n%s\n", prettyA, prettyB, dmp.DiffPrettyText(diffs))
}

// ExpectedGotString returns a formatted string for
// the normal use case of "expected a got b" string
func ExpectedGotString(a string, b string) string {
	return fmt.Sprintf("\nexpected: \"%s\" got: \"%s\"\n", a, b)
}
