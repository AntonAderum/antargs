# antargs

An argument package for go.

Basic usage:
```go
// setup
antArg, _ := New("myTestProgram", "myTestProgram is a very nice program!")
antArg.NewArg("--say", "What should I say?", false, "", 1)
antArg.NewArg("--style", "In what style should I say it?", false, "", 1)

// parse
antArg.Parse(os.Args)

// use
sayArg := antArg.findArgument("--say")
sayValue := arg.values[0]
styleValue := antArg.findArgument("--style").values[0]

if styleValue == "shout" {
    sayValue = ToUpper(sayValue)
}
fmt.Printf("I say: %s", sayValue)
```
terminal:
```bash
$ ./myTestProgram --say "You're A Wizard, Harry!" --style shout
$ I say: YOU'RE A WIZARD, HARRY!
```


