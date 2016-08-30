package envflag

import (
	"flag"
	"os"
	"testing"
)

func setupFlag(options flagOption) *Envflag {
	os.Clearenv()
	cli := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	options(cli)
	return &Envflag{
		Cli: cli,
	}
}

// declare one flag, assigns value via cli, expects empty unsetFlags
func TestOneSetFlag(t *testing.T) {
	var testVal string

	ef := setupFlag(func(fs *flag.FlagSet) {
		fs.StringVar(&testVal, "test-flag", "default-value", "this is test flag")
	})

	ef.parseWithEnv([]string{"--test-flag=dummy"})

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
		t.Errorf("expects only %v unset flag, got %v.", want, got)
	}
}

// default < env < cli
func TestPrecendence(t *testing.T) {

	var (
		testVal    string
		envVal     = "envValue"
		cliVal     = "cliValue"
		defaultVal = "defaultValue"
		// testIntTwo int
	)

	// cli
	ef := setupFlag(func(fs *flag.FlagSet) {
		fs.StringVar(&testVal, "test-flag", defaultVal, "this is test flag")
	})

	os.Setenv("TEST_FLAG", envVal)
	ef.parseWithEnv([]string{
		"--test-flag=" + cliVal,
	})

	if want, got := cliVal, testVal; want != got {
		t.Errorf("expects flag-value be overriden by (cli)%v, got %v.", want, got)
	}

	// env
	ef = setupFlag(func(fs *flag.FlagSet) {
		fs.StringVar(&testVal, "test-flag", defaultVal, "this is test flag")
	})
	os.Setenv("TEST_FLAG", envVal)

	ef.parseWithEnv([]string{}) // empty cli
	if want, got := envVal, testVal; want != got {
		t.Errorf("expects flag-value be overriden by (env)%v, got %v.", want, got)
	}

	// default
	ef = setupFlag(func(fs *flag.FlagSet) {
		fs.StringVar(&testVal, "test-flag", defaultVal, "this is test flag")
	})

	ef.parseWithEnv([]string{}) // empty cli
	if want, got := defaultVal, testVal; want != got {
		t.Errorf("expects flag-value be overriden by (env)%v, got %v.", want, got)
	}
}

// flagOption for testing purposes -- visibility
type flagOption func(*flag.FlagSet)
