package txtfsm

type Funcs struct {
	FuncName  string
	Func      *FSMInterface
	Params    FuncParams
	Signature FuncSignature
}

type FuncSignature struct {
}

func (fs *FuncSignature) ValidateSignature() bool {
	return false
}
