package txtfsm

import "sync"

type FSMVar int

const (
	VInt FSMVar = iota
	VFloat
	VString
)

type Var struct {
	Name     string
	VType    FSMVar
	noUpdate bool
	lock     *sync.Mutex
	curval   interface{}
}

func (v *Var) AsInt() int {
	return v.curval.(int)
}

func (v *Var) AsFloat() float32 {
	return v.curval.(float32)
}

func (v *Var) AsString() string {
	return v.curval.(string)
}

func (v *Var) SetVal(nv interface{}) {
	if v.noUpdate {
		return
	}
	v.lock.Lock()
	defer v.lock.Unlock()
	switch v.VType {
	case VInt:
		v.curval = nv.(int)
	case VFloat:
		v.curval = nv.(float32)
	case VString:
		v.curval = nv.(string)
	}
}

func NewVar(nu bool) *Var {
	return &Var{lock: &sync.Mutex{}, noUpdate: nu}
}
