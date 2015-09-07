package tomita

import (
	"io/ioutil"
	"os"
	"reflect"
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

	shouldBe := Result{Facts: map[string][]map[string]string{"Group": []map[string]string{map[string]string{"Name": "ГРУППА LOUNA"}}, "Album": []map[string]string{map[string]string{"Name": "АЛЬБОМА  МЫ — ЭТО LOUNA !"}, map[string]string{"Name": "АЛЬБОМ  ПРОСНИСЬ И ПОЙ"}}}, Leads: []string{"4 апреля в клубе \"РОК-СИТИ\"  даст большой сольный концерт, приуроченный к выходу своего нового , над которым музыканты сейчас работают в студии.", "В этом году музыканты успели выпустить концертный DVD и живой , с размахом провести его презентацию в клубе \"ARENA Moscow\", съездить в масштабный тур по России и СНГ и выступить на двадцати летних фестивалях."}}
	if !reflect.DeepEqual(out, shouldBe) {
		t.Fatalf("unexpected reply:\n %v \nshould be: \n %v\n", out, shouldBe)
	}

	t.Log("the facts are: ")
	for factsGroupName, facts := range out.Facts {
		for _, values := range facts {
			for factName, val := range values {
				t.Log(factsGroupName + ": " + factName + ": " + val)
			}
		}
	}

	t.Log("the leads are: ")
	for _, lead := range out.Leads {
		t.Log(lead)
	}
}

func fatalOnErr(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
