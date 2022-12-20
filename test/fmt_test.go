package test

import (
	"github.com/995933447/stringhelper-go"
	"testing"
)

func TestCamel(t *testing.T) {
	t.Log(stringhelper.Camel("Ab_cd_ef"))
}
