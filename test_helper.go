package antargs

import "fmt"

var (
	arguments_shortcut    = []string{"", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	subarguments_shortcut = []string{"", "aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj", "kk", "ll", "mm", "nn", "oo", "pp", "qq", "rr", "ss", "tt", "uu", "vv", "ww", "xx", "yy", "zz"}
)

func getTestingAntArgObject() *AntArg {
	return &AntArg{
		name: "test",
		help: "help_test",
		args: []*Arg{},
	}
}

func getTestingAntArgObjectWithArgument(arguments int) *AntArg {
	antArg := &AntArg{
		name: "test",
		help: "help_test",
	}
	var args []*Arg

	for i := 0; i < arguments; i++ {
		args = append(args, &Arg{
			help:           fmt.Sprintf("sub_help_%d", i),
			name:           fmt.Sprintf("sub_name_%d", i),
			isFlag:         i%2 == 1,
			shortcut:       arguments_shortcut[i],
			numberOfValues: 1,
			values:         []string{},
		})
	}
	antArg.args = args
	return antArg
}

func getTestingAntArgObjectWithArgumentAndSubArguments(arguments int, subArguments int) *AntArg {
	antArg := &AntArg{
		name: "test",
		help: "help_test",
	}
	var args []*Arg
	for i := 0; i < arguments; i++ {
		var subArgs []*Arg
		for j := 0; j < subArguments; j++ {
			subArgs = append(subArgs, &Arg{
				name:           fmt.Sprintf("sub_sub_name_%d", j),
				help:           fmt.Sprintf("sub_sub_help_%d", j),
				isFlag:         i%2 == 1,
				shortcut:       subarguments_shortcut[j],
				numberOfValues: 1,
				values:         []string{},
			})
		}
		args = append(args, &Arg{
			help:           fmt.Sprintf("sub_help_%d", i),
			name:           fmt.Sprintf("sub_name_%d", i),
			isFlag:         i%2 == 1,
			shortcut:       arguments_shortcut[i],
			subArgs:        subArgs,
			numberOfValues: 1,
			values:         []string{},
		})
	}
	antArg.args = args
	return antArg
}

func getTestingAntArgObjectWithArgumentAndSubArgumentsAndSubArguments(arguments int, subArguments int, subSubArguments int) *AntArg {
	antArg := &AntArg{
		name: "test",
		help: "help_test",
	}
	var args []*Arg
	for i := 0; i < arguments; i++ {
		var subArgs []*Arg
		for j := 0; j < subArguments; j++ {
			var subSubArgs []*Arg
			for k := 0; k < subSubArguments; k++ {
				subSubArgs = append(subSubArgs, &Arg{
					name:           fmt.Sprintf("sub_sub_sub_name_%d", k),
					help:           fmt.Sprintf("sub_sub_sub_help_%d", k),
					isFlag:         i%2 == 1,
					shortcut:       "",
					numberOfValues: 1,
					values:         []string{},
				})
			}
			subArgs = append(subArgs, &Arg{
				name:           fmt.Sprintf("sub_sub_name_%d", j),
				help:           fmt.Sprintf("sub_sub_help_%d", j),
				isFlag:         i%2 == 1,
				shortcut:       subarguments_shortcut[j],
				numberOfValues: 1,
				subArgs:        subSubArgs,
				values:         []string{},
			})
		}
		args = append(args, &Arg{
			help:           fmt.Sprintf("sub_help_%d", i),
			name:           fmt.Sprintf("sub_name_%d", i),
			isFlag:         i%2 == 1,
			shortcut:       arguments_shortcut[i],
			subArgs:        subArgs,
			numberOfValues: 1,
			values:         []string{},
		})
	}
	antArg.args = args
	return antArg
}
