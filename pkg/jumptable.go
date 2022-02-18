package txtfsm

const MaxStackSize = 20 //Prevent infinite recursion

type JumpElement struct {
	Current interface{}
	Next    interface{}
}

type JumpTable struct {
	InitState     *JumpElement
	StateSequence []*JumpElement
}

type StateStack struct {
	Element []JumpElement
}

type TxtFSMRegistry struct {
	

}
