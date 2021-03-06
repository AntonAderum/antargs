package antargs

import (
	"fmt"
)

type parseOption int32

const (
	allowOnlyOneTopLevelArgument      parseOption = 1
	dontRemoveFirstArgument           parseOption = 2
	requireAtleastOneTopLevelArgument parseOption = 3
	allowTopLevelArgumentWithoutName  parseOption = 4
)

func AllowOnlyOneTopLevelArgument() parseOption      { return allowOnlyOneTopLevelArgument }
func DontRemoveFirstArgument() parseOption           { return dontRemoveFirstArgument }
func RequireAtleastOneTopLevelArgument() parseOption { return requireAtleastOneTopLevelArgument }
func AllowTopLevelArgumentWithoutName() parseOption  { return allowTopLevelArgumentWithoutName }

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

func findArgumentFromArray(args []*Arg, name string) *Arg {
	// This is a very naive implementation take could take a long time
	// with lots of arguments. Maybe in the future it could benefit from
	// some more clever search logic, but for now I think its  okay
	for _, arg := range args {
		if arg.name == name {
			return arg
		}
	}
	return nil
}
func (antArg *AntArg) findArgument(name string) *Arg {
	return findArgumentFromArray(antArg.args, name)
}
func (arg *Arg) findSubArgument(name string) *Arg {
	return findArgumentFromArray(arg.subArgs, name)
}

// Parse takes an array of string arguments assigns the arguments to their
// corresponding value in the AntArg object.
//
// parseOptions can be provided to control the parsing:
// AllowOnlyOneTopLevelArgument() = Will return error if more than 1 argument is provided at the top level (AntArg.args)
// DontRemoveFirstArgument() = Will disable the default behavior to remove the first element of the argument array
// RequireAtleastOneTopLevelArgument() = Will return error if no argument has been provided
// AllowTopLevelArgumentWithoutName() = Will allow the top level arguments to be provided with only values, this will assume no top level argument is provided with name, and that the values are provided in the order that the arguments was provided into the AntArgs object
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
	// ToDo: Handle allowing top level arguments without supplying name for them
	state := parseState(start)
	var currentArg *Arg
	var currentSubArg *Arg
	currentArgIndex := 0

	if parseOptionWasRequested(allowTopLevelArgumentWithoutName, parseOptions) {
		state = parseState(readingArgValues)
		currentArg = antArg.args[currentArgIndex]
	}
	for i, argument := range arguments {
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
			loops := -1
			for {
				loops = loops + 1
				if len(arguments) < (i + loops) {
					break
				}
				// ToDo: I think this break with something like:
				// ./test-program argument --subArgument 1 --subArgument a
				currentSubArg = currentArg.findSubArgument(arguments[i+loops])
				if currentSubArg == nil {
					state = parseState(readingArgValues)
					break
				} else if currentSubArg.isFlag {
					// A sub argument was provided and was a flag
					// keep state in foundArg to see if more sub arguments was provided
					currentSubArg.wasProvided = true
				} else {
					state = parseState(readingSubArgValues)
					break
				}
			}
		} else if state == readingArgValues {
			currentArgLength := len(currentArg.values)
			if currentArgLength < currentArg.numberOfValues {
				currentArg.values = append(currentArg.values, argument)
				currentArgLength = currentArgLength + 1
			}
			if currentArgLength == currentArg.numberOfValues {
				currentArg.wasProvided = true
				if parseOptionWasRequested(allowTopLevelArgumentWithoutName, parseOptions) {
					if currentArgIndex < len(antArg.args)-1 {
						currentArgIndex = currentArgIndex + 1
						currentArg = antArg.args[currentArgIndex]
					}
				} else {
					state = parseState(start)
				}
			}
		} else if state == readingSubArgValues {
			// ToDo: state == readingSubArgValues and state == readingArgValues
			// are doing exactly the same thing, make this into one method/logic block?
			currentSubArgLength := len(currentSubArg.values)
			if currentSubArgLength < currentSubArg.numberOfValues {
				currentSubArg.values = append(currentSubArg.values, argument)
				currentSubArgLength = currentSubArgLength + 1
			}
			if currentSubArgLength == currentSubArg.numberOfValues {
				currentSubArg.wasProvided = true
				state = parseState(readingArgValues)
			}
		}
	}
	if state == readingArgValues {
		return fmt.Errorf("Not enough values was provided. %s expects %d values and got %d\n", currentArg.name, currentArg.numberOfValues, len(currentArg.values))
	}
	if state == readingSubArgValues {
		return fmt.Errorf("Not enough values was provided. %s expects %d values and got %d\n", currentSubArg.name, currentSubArg.numberOfValues, len(currentSubArg.values))
	}
	return nil
}
