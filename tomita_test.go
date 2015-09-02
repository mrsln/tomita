package tomita

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	p, err := New("/bin/tomita", "example/config.proto")
	fatalOnErr(err, t)

	wd, _ := os.Getwd()
	if p.originalWd != wd {
		t.Fatal("wrong original working directory: " + p.originalWd)
	}

	txt, err := ioutil.ReadFile("example/text.txt")
	fatalOnErr(err, t)

	out, err := p.Run(string(txt))
	fatalOnErr(err, t)

	if len(out.Leads) == 0 || out.Facts == nil {
		t.Fatalf("the parser didn't parse anything: %#v", out)
	}
	t.Log("the facts are: ")
	for factsGroupName, facts := range out.Facts {
		for _, values := range facts {
			for factName, val := range values {
				t.Log(factsGroupName + ": " + factName + ": " + val)
			}
		}
	}
}

func fatalOnErr(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
