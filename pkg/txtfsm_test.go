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

func (ms *myState) pre_run(params []interface{}) (bool, error) {
	return true, nil
}

func (ms *myState) post_run(params []interface{}) (bool, error) {
	return true, nil
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

func (ms *myNextState) pre_run(params []interface{}) (bool, error) {
	return true, nil
}

func (ms *myNextState) post_run(params []interface{}) (bool, error) {
	return true, nil
}

func TestFSMRegisterSingleFunc(t *testing.T) {
	SetTestMode()
	TxtFSMClearMap()
	ms := &myState{}
	fp := make(FuncParams, 0)
	if !TxtFSMRegister("TestState1", ms, fp) {
		t.Fail()

	}

}

func TestFSMRegisterDuplicateFunc(t *testing.T) {
	SetTestMode()
	TxtFSMClearMap()
	ms := &myState{}
	fp := make(FuncParams, 0)
	if !TxtFSMRegister("TestState1", ms, fp) {
		t.Logf("Already registered TestState1\n")
		t.Fail()
	}
	ns := &myNextState{}
	if TxtFSMRegister("TestState1", ns, fp) {
		t.Logf("Should have already registered TestState1\n")
		t.Fail()
	}

}

func TestFSMRegisterMultipleFunc(t *testing.T) {
	SetTestMode()
	TxtFSMClearMap()
	ms := &myState{}
	fp := make(FuncParams, 0)
	if !TxtFSMRegister("S1Func1", ms, fp) {
		t.Logf("Already registered S1Func1\n")
		t.Fail()

	}
	ns := &myNextState{}
	np := make(FuncParams, 0)
	np = np.Params(FSMConstant, FSMInt).Params(FSMConstant, FSMString)
	fmt.Printf("%v %v\n", *np[0], *np[1])
	if !TxtFSMRegister("S1Func2", ns, np) {
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
	fp := make(FuncParams, 0)

	if !TxtFSMRegister("S1Func1", ms, fp) {
		t.Logf("Already registered S1Func1\n")
		t.Fail()

	}
	ns := &myNextState{}
	np := make(FuncParams, 0)
	np = np.Params(FSMConstant, FSMInt).Params(FSMConstant, FSMString)

	if !TxtFSMRegister("S1Func2", ns, np) {
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
