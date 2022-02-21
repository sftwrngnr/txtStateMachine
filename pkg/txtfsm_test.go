package txtfsm

import (
	"errors"
	"fmt"
	"testing"
)

// Test Functions

type myState struct {
}

func (ms *myState) run(params []interface{}) (int, error) {
	fmt.Printf("Test State 1 %v\n", params)
	return 0, nil
}

type myNextState struct {
}

func (mns *myNextState) run(params []interface{}) (int, error) {
	fmt.Printf("Test State 2 %v\n", params)
	if params[0].(int) != 3 {
		return 0, errors.New("incorrect parameter passed")
	}
	return 1, nil
}

func TestFSMRegisterSingleFunc(t *testing.T) {
	ms := &myState{}
	if !TxtFSMRegister("TestState1", ms) {
		t.Fail()

	}

}

func TestFSMRegisterDuplicateFunc(t *testing.T) {
	TxtFSMClearMap()
	ms := &myState{}
	if !TxtFSMRegister("TestState1", ms) {
		t.Logf("Already registered TestState1\n")
		t.Fail()
	}
	ns := &myNextState{}
	if TxtFSMRegister("TestState1", ns) {
		t.Logf("Already registered TestState1\n")
		t.Fail()
	}

}

func TestFSMRegisterMultipleFunc(t *testing.T) {
	TxtFSMClearMap()
	ms := &myState{}
	if !TxtFSMRegister("S1Func1", ms) {
		t.Logf("Already registered S1Func1\n")
		t.Fail()

	}
	ns := &myNextState{}
	if !TxtFSMRegister("S1Func2", ns) {
		t.Logf("Already registered S1Func2\n")
		t.Fail()
	}
	if _, err := TestRun("S1Func1", nil); err != nil {
		t.Logf("%s", err.Error())
		t.Fail()
	}
	paramsa := make([]interface{}, 0)
	paramsa = append(paramsa, 3)
	paramsa = append(paramsa, "Blah")
	if _, err := TestRun("S1Func2", paramsa); err != nil {
		t.Logf("%s", err.Error())
		t.Fail()
	}
	paramsb := make([]interface{}, 0)
	paramsb = append(paramsb, 2)
	paramsb = append(paramsb, "Blah")
	if _, err := TestRun("S1Func2", paramsb); err == nil {
		t.Logf("Should have thrown an error!\n")
		t.Fail()
	}

}

func TestInvalidFuncName(t *testing.T) {
	TxtFSMClearMap()
	ms := &myState{}
	if !TxtFSMRegister("S1Func1", ms) {
		t.Logf("Already registered S1Func1\n")
		t.Fail()

	}
	ns := &myNextState{}
	if !TxtFSMRegister("S1Func2", ns) {
		t.Logf("Already registered S1Func2\n")
		t.Fail()
	}
	if _, err := TestRun("S1Func3", nil); err == nil {
		t.Logf("Should have thrown an error!\n")
		t.Fail()
	} else {
		fmt.Printf("%s\n", err.Error())
	}

}
