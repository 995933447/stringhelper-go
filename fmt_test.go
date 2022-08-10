package stringhelper

import (
	"testing"
)

func TestCamel(t *testing.T) {
	t.Log(Camel("Ab_cd_ef"))
}
