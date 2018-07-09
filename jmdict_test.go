package jmdict

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	file, err := os.Open("./data/JMdict.xml")
	if err != nil {
		t.Error(err)
	}
	dict := Load(file)
	if dict == nil {
		t.Error("dict is nil!\n")
	} else if dict.Entries == nil {
		t.Error("no entries in dict!\n")
	}
	t.Logf("%v\n", dict.Entries[20])
}
