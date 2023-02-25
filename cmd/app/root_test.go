package cmd_test

import (
	"testing"

	cmd "github.com/asosnoviy/go1cover/cmd/app"
	"github.com/matryer/is"
)

func Test_root(t *testing.T) {

	is := is.New(t)
	err := cmd.Execute()

	is.NoErr(err)
}
