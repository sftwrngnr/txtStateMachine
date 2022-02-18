package txtfsm

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
	switch v.VType {
	case VInt:
		v.curval = nv.(int)
	case VFloat:
		v.curval = nv.(float32)
	case VString:
		v.curval = nv.(string)
	}
}
