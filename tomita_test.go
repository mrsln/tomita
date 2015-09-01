package tomita

import (
	"io/ioutil"
	"testing"
)

func TestRun(t *testing.T) {
	txt, err := ioutil.ReadFile("example/text.txt")
	fatalOnErr(err, t)

	p, err := New("/bin/tomita", "example/config.proto")
	fatalOnErr(err, t)

	str, err := p.Run(string(txt))
	fatalOnErr(err, t)

	t.Log(str)

	if len(str) == 0 {
		t.Fatal("the parser didn't parse anything")
	}
}

func fatalOnErr(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
