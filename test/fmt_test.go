package test

import (
	"testing"

	"github.com/995933447/stringhelper-go"
)

func TestFmt(t *testing.T) {
	t.Log(stringhelper.Camel("Ab_cd_ef"))
	t.Log(stringhelper.Snake("Ab_cd_ef"))
}
