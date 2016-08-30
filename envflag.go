package envflag

import (
	"flag"
	"os"
	"strings"
)

// Envflag is a environment-variable wrapper for flag.FlagSet.
type Envflag struct {
	Cli *flag.FlagSet
}

// DefaultEnvflag is the default Envflag and is used by Parse.
var DefaultEnvflag = &Envflag{
	Cli: flag.CommandLine,
}

// Parse parses the command-line flags from os.Args[1:].
// For unset flags, it will attempt to read & set from environment variables.
// Must be called after all flags are defined and before flags are accessed by the program.
func (e Envflag) Parse() {
	e.parseWithEnv(os.Args[1:])
}

// parseWithEnv first parses cli-values into flag-values. Next, for any unset
// flag-values, it attemps to lookup and load from environment variables, by flag-name.
func (e Envflag) parseWithEnv(args []string) {

	if !e.Cli.Parsed() {
		e.Cli.Parse(args)
	}

	for name := range e.unsetFlags() {
		if val, ok := os.LookupEnv(maskEnvName(name)); ok {
			e.Cli.Set(name, val)
		}
	}
}

// maskEnvName returns a qualified environement variable, naming convention.
func maskEnvName(flagname string) string {
	var ret string

	ret = strings.Replace(flagname, "-", "_", -1)
	ret = strings.Replace(flagname, ".", "_", -1)
	return strings.ToUpper(ret)
}

// unsetFlags returns flag-values that hasn't been set via CLI.
func (e Envflag) unsetFlags() map[string]struct{} {

	flags := make(map[string]struct{})

	// get all flag-names.
	e.Cli.VisitAll(func(f *flag.Flag) {
		flags[f.Name] = struct{}{}
	})

	// MINUS from set flag-names.
	e.Cli.Visit(func(f *flag.Flag) {
		delete(flags, f.Name)
	})

	return flags
}

// Parse parses the command-line flags from os.Args[1:].
// Must be called after all flags are defined and before flags are accessed by the program.
func Parse() {
	DefaultEnvflag.parseWithEnv(os.Args[1:])
}
