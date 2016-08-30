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

// set one flag via cli, expects empty unsetFlags
func TestOneSetFlag(t *testing.T) {
	var testInt int

	ef := setupFlag(func(fs *flag.FlagSet) {
		fs.IntVar(&testInt, "test-int", 1, "this is test int")
	})

	ef.parseWithEnv([]string{"--test-int=100"})

	if want, got := 0, len(ef.unsetFlags()); want != got {
		t.Errorf("non-empty unsetflags. (want)%v != %v(got)", want, got)
	}
}

type flagOption func(*flag.FlagSet)
