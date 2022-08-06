package stringhelper

import (
	"testing"
)

func TestCamel(t *testing.T) {
	t.Log(Camel("a_bc_d"))
}
