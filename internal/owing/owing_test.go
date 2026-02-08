package owing_test

import (
	"testing"

	"github.com/matteo-gildone/owing/internal/owing"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"owing": owing.Main,
	})
}

func Test(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "../../testdata/script",
	})
}
