package txtfsm

type FSMFunc struct {
	Name      string
	Func      *FSMInterface
	Params    FuncParams
	Signature FuncSignature
}

type FSMFuncList []*FSMFunc

type FuncSignature struct {
}

func (fs *FuncSignature) ValidateSignature() bool {
	return false
}
