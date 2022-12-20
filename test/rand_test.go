package test

import (
	"github.com/995933447/stringhelper-go"
	"testing"
)

func TestRandomStr(t *testing.T) {
	t.Log(stringhelper.GenRandomStr(6, stringhelper.RandomStringModNumberPlusLetterPlusSymbol))
	t.Log(stringhelper.GenRandomStr(6, stringhelper.RandomStringModNumberPlusLetterPlusSymbol))
}
