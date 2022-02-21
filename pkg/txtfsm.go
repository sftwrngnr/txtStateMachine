package txtfsm

import (
	"errors"
	"fmt"
	"sync"
)

type TxtFSM struct {
	fMap map[string]TxtFSMFunc
}

type TxtFSMFunc interface {
	run(params ...interface{}) (int, error)
}

// Singleton var
var sTextFSM *TxtFSM
var stfLock = &sync.Mutex{}
var addLock = &sync.Mutex{}

func getInstance() *TxtFSM {
	if sTextFSM == nil {
		stfLock.Lock()
		defer stfLock.Unlock()
		sTextFSM = &TxtFSM{fMap: make(map[string]TxtFSMFunc)}
	}
	return sTextFSM
}

func (t *TxtFSM) addFunc(name string, f TxtFSMFunc) bool {
	// Protect the instance with a mutext while adding
	addLock.Lock()
	defer addLock.Unlock()
	if _, ok := t.fMap[name]; !ok {
		t.fMap[name] = f
		return true
	}
	return false
}

func (t *TxtFSM) testRun(fName string, params ...interface{}) (int, error) {
	if _, ok := t.fMap[fName]; !ok {
		return 0, errors.New("invalid function called")

	}
	fmt.Printf("Running %s\n", fName)
	return t.fMap[fName].run(params)

}

func TestRun(fname string, params ...interface{}) (int, error) {
	return getInstance().testRun(fname, params)
}

func TxtFSMRegister(name string, f TxtFSMFunc) bool {
	if !getInstance().addFunc(name, f) {
		fmt.Printf("Function %s already registered!!\n", name)
		return false
	}
	return true
}

func (t *TxtFSM) clearMap() {
	// This should only be called if we are testing.
	addLock.Lock()
	defer addLock.Unlock()
	t.fMap = make(map[string]TxtFSMFunc)
}

func TxtFSMClearMap() {
	getInstance().clearMap()
}
