package antargs

// Equal compares two AntArg objects and decied if they are equal
func (a AntArg) Equal(b AntArg) bool {
	return a.help == b.help &&
		a.name == b.name &&
		deepCompareArgs(a.args, b.args)
}

func compareValues(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func deepCompareArgs(a []*Arg, b []*Arg) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i].name != b[i].name ||
			a[i].help != b[i].help ||
			a[i].shortcut != b[i].shortcut ||
			!compareValues(a[i].values, b[i].values) ||
			!deepCompareArgs(a[i].subArgs, b[i].subArgs) {
			return false
		}
	}
	return true
}
