package txtfsm

import (
	"fmt"
	"sync"
)

type txtFSMElement struct {
	fint   FSMInterface
	fparms FuncParams
}

type TxtFSM struct {
	fMap     map[string]txtFSMElement
	testMode bool
	addLock  *sync.Mutex
}

// Singleton var
var sTextFSM *TxtFSM
var stfLock = &sync.Mutex{}

func getInstance() *TxtFSM {
	if sTextFSM == nil {
		stfLock.Lock()
		defer stfLock.Unlock()
		sTextFSM = &TxtFSM{fMap: make(map[string]txtFSMElement), addLock: &sync.Mutex{}}
	}
	return sTextFSM
}

func (t *TxtFSM) addFunc(name string, f FSMInterface, p FuncParams) bool {
	// Protect the instance with a mutext while adding
	t.addLock.Lock()
	if t.testMode {
		fmt.Printf("Adding function %s", name)
	}
	defer t.addLock.Unlock()
	if _, ok := t.fMap[name]; !ok {
		t.fMap[name] = txtFSMElement{fint: f, fparms: p}
		if t.testMode {
			fmt.Printf(" successfully\n")
		}
		return true
	}
	if t.testMode {
		fmt.Printf(" unsuccessfully\n")
	}
	return false
}

func (t *TxtFSM) testRun(fName string, params []interface{}) (int, error) {
	if _, ok := t.fMap[fName]; !ok {
		return 0, fmt.Errorf("invalid function [%s] called", fName)

	}
	if t.testMode {
		fmt.Printf("Running %s\n", fName)
	}
	return t.fMap[fName].fint.run(params)

}

func TestRun(fname string, params []interface{}) (int, error) {
	return getInstance().testRun(fname, params)
}

func TxtFSMRegister(name string, f FSMInterface, p FuncParams) bool {
	return getInstance().addFunc(name, f, p)
}

func (t *TxtFSM) clearMap() {
	// This should only be called if we are testing.
	if !t.testMode {
		return
	}
	fmt.Printf("clearing function map\n")
	t.addLock.Lock()
	defer t.addLock.Unlock()
	t.fMap = make(map[string]txtFSMElement)
}

func TxtFSMClearMap() {
	getInstance().clearMap()
}

func (t *TxtFSM) setTestMode() {
	t.addLock.Lock()
	defer t.addLock.Unlock()
	t.testMode = true
}

func SetTestMode() {
	getInstance().setTestMode()
}
