package envflag

import (
	"flag"
	"os"
	"testing"
)

func setupFlag(options flagOption) *Envflag {
	cli := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	options(cli)
	return &Envflag{
		Cli: cli,
	}
}

// declare one flag, assigns value via cli, expects empty unsetFlags
func TestOneSetFlag(t *testing.T) {
	var testInt int

	ef := setupFlag(func(fs *flag.FlagSet) {
		fs.IntVar(&testInt, "test-int", 1, "this is test int")
	})

	ef.parseWithEnv([]string{"--test-int=100"})

	if want, got := 0, len(ef.unsetFlags()); want != got {
		t.Errorf("expects only %v unset flag, got %v intead.", want, got)
	}
}

// declare two flags, only set 1 via cli, expects 1 unsetflag
func TestTwoFlagsOneUnset(t *testing.T) {
	var (
		testIntOne int
		testIntTwo int
	)

	ef := setupFlag(func(fs *flag.FlagSet) {
		fs.IntVar(&testIntOne, "test-int-one", 1, "this is test int one")
		fs.IntVar(&testIntTwo, "test-int-two", 1, "this is test int two")
	})

	ef.parseWithEnv([]string{"--test-int-one=100"})

	if want, got := 1, len(ef.unsetFlags()); want != got {
		t.Errorf("expects only %v unset flag, got %v intead.", want, got)
	}
}

// flagOption for testing purposes -- visibility
type flagOption func(*flag.FlagSet)
