package txtfsm

import "testing"

func TestParserRead(t *testing.T) {
	fp := &FSMParser{}
	if err := fp.ReadFile("./testdata/testdata.fsm"); err != nil {
		t.Fail()
	}
	fp.Parse()
}
