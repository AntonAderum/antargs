package antargs

import (
	"fmt"
)

type parseOption int32

const (
	allowOnlyOneTopLevelArgument      parseOption = 1
	dontRemoveFirstArgument           parseOption = 2
	requireAtleastOneTopLevelArgument parseOption = 3
)

func AllowOnlyOneTopLevelArgument() parseOption      { return allowOnlyOneTopLevelArgument }
func DontRemoveFirstArgument() parseOption           { return dontRemoveFirstArgument }
func RequireAtleastOneTopLevelArgument() parseOption { return requireAtleastOneTopLevelArgument }

type parseState int32

const (
	start               parseState = 1
	foundArg            parseState = 2
	readingArgValues    parseState = 3
	foundSubArg         parseState = 4
	readingSubArgValues parseState = 5
)

func parseOptionWasRequested(option parseOption, options []parseOption) bool {
	for _, opt := range options {
		if opt == option {
			return true
		}
	}
	return false
}

func (antArg *AntArg) findArgument(name string) *Arg {
	for _, arg := range antArg.args {
		if arg.name == name {
			return arg
		}
	}
	return nil
}

// Parse takes an array of string arguments assigns the arguments to their
// corresponding value in the AntArg object.
//
// parseOptions can be provided to control the parsing:
// AllowOnlyOneTopLevelArgument() = Will return error if more than 1 argument is provided at the top level (AntArg.args)
// DontRemoveFirstArgument() = The default behavior is to remove the first element of the argument array, this settings disables that behavior
// RequireAtleastOneTopLevelArgument() = Will return error if no argument has been provided
//
// The following errors can be returned:
// "No argument supplied" = Will happen if RequireOneTopLevelArgument parseOption was requested and no argument was provided
// "%s it not a valid argument" = Will happen if a top level argument was supplied but no corresponding argument was declared in the AntAr
// "Only one top level argument is allowed" == Will happening if more than 1 top level argument was provided and the parseOption AllowonlyOneTopLevelArgument was request
func (antArg *AntArg) Parse(arguments []string, parseOptions ...parseOption) error {
	if !parseOptionWasRequested(dontRemoveFirstArgument, parseOptions) {
		arguments = arguments[1:]
	}
	if parseOptionWasRequested(requireAtleastOneTopLevelArgument, parseOptions) && len(arguments) == 0 {
		return fmt.Errorf("No argument supplied")
	}
	// This is a very naive implemention of this, and probably very slow if
	// lots of arguments are being parsed / in the AntArg.args list.
	// This should probably/maybe be written in some more clever way,
	// but it's now written to "just work"™
	state := parseState(start)
	var currentArg *Arg
	for _, argument := range arguments {
		if state == start {
			if currentArg != nil && parseOptionWasRequested(allowOnlyOneTopLevelArgument, parseOptions) {
				return fmt.Errorf("Only one top level argument is allowed")
			}
			currentArg = antArg.findArgument(argument)
			if currentArg == nil {
				return fmt.Errorf("%s is not a valid argument", argument)
			}
			if len(currentArg.subArgs) == 0 {
				// If arguments doens't have sub arguments
				// we go straight to reading values
				state = parseState(readingArgValues)
			} else if currentArg.isFlag {
				currentArg.wasProvided = true
				state = parseState(start)
			} else {
				state = parseState(foundArg)
			}
		} else if state == foundArg {
			// ToDo: parse sub args
			state = state
		} else if state == readingArgValues {
			currentArgLength := len(currentArg.values)
			if currentArgLength < currentArg.numberOfValues {
				currentArg.values = append(currentArg.values, argument)
			}
			if currentArgLength < currentArg.numberOfValues {
				state = parseState(start)
			}
		}
	}
	if state == readingArgValues {
		return fmt.Errorf("Not enough values was provided. %s expects %d values", currentArg.name, currentArg.numberOfValues)
	}
	return nil
}
