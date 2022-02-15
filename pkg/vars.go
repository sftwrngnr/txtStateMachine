package txtfsm

type FSMVar int

const (
	VInt FSMVar = iota
	VFloat
	VString
)

type Vars struct {
	Name   string
	VType  FSMVar
	CurVal interface{}
}

func (v *Vars) AsInt() int {
	return v.CurVal.(int)
}

func (v *Vars) AsFloat() float32 {
	return v.CurVal.(float32)
}

func (v *Vars) AsString() string {
	return v.CurVal.(string)
}
