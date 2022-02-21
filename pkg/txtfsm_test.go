package txtfsm

import (
	"fmt"
	"testing"
)

// Test Functions

type myState struct {
}

func (ms *myState) run(params ...interface{}) (int, error) {
	fmt.Printf("Test State 1 %v\n", params)
	return 0, nil
}

type myNextState struct {
}

func (mns *myNextState) run(params ...interface{}) (int, error) {
	fmt.Printf("Test State 2 %v\n", params)
	return 1, nil
}

func TestFSMRegisterSingleFunc(t *testing.T) {
	ms := &myState{}
	if !TxtFSMRegister("TestState1", ms) {
		t.Fail()

	}

}

func TestFSMRegisterDuplicateFunc(t *testing.T) {
	ms := &myState{}
	if !TxtFSMRegister("TestState1", ms) {
		t.Fail()

	}
	ns := &myNextState{}
	if TxtFSMRegister("TestState1", ns) {
		t.Fail()
	}

}

func TestFSMRegisterMultipleFunc(t *testing.T) {
	ms := &myState{}
	if !TxtFSMRegister("TestState1", ms) {
		t.Fail()

	}
	ns := &myNextState{}
	if !TxtFSMRegister("TestState2", ns) {
		t.Fail()
	}
	if _, err := TestRun("TestState1", nil); err != nil {
		t.Logf("%s", err.Error())
		t.Fail()
	}
	if _, err := TestRun("TestState2", 3, "Blah"); err != nil {
		t.Logf("%s", err.Error())
	}

}
