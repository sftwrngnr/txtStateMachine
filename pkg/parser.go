package txtfsm

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getReservedList() []string {
	return []string{":EXEC", ":DONE", ":RETURN", ":CONTINUE", ":VARS", ":ABORT", ":RESTART"}
}

func getTerminusList() []string {
	return []string{":RETURN", ":CONTINUE", ":ABORT", ":RESTART", ":DONE"}
}

type FSMElements struct {
	FSMLabels FSMLabelMap
	FSMFuncs  FSMFuncList
	FSMVars   []Var
}

type FSMParser struct {
	FsmText []string
	Pmap    FSMElements
}

func (f *FSMParser) ReadFile(fn string) error {
	data, err := os.ReadFile(fn)
	if err != nil {
		return err
	}
	f.FsmText = strings.Split(string(data), "\n")
	return nil
}

func (f *FSMParser) trimComment(s string) string {
	if strings.Contains(s, "#") {
		return s[0:strings.Index(s, "#")]
	}
	return s
}

func (f *FSMParser) checkReserved(s string) bool {
	for _, v := range getReservedList() {
		if strings.Contains(s, v) {
			fmt.Printf("Reserved word: %s\n", v)
			return true
		}
	}
	return false
}

func (f *FSMParser) checkTerminus(s string) bool {
	for _, v := range getTerminusList() {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}

func (f *FSMParser) getVarName(v, s string) string {
	retval := s[0:strings.Index(s, v)]
	return strings.Trim(retval, " \t")
}

func (f *FSMParser) getVarIntVal(s string) (int, error) {
	vIdx := strings.Index(s, "INT")
	if vIdx > -1 {
		vIdx += 3
		cval := strings.Trim(s[vIdx:], " \t\n")
		if len(cval) == 0 {
			return 0, nil
		}
		rv, err := strconv.Atoi(cval)
		if err != nil {
			return 0, err // Failed to parse
		}
		return rv, nil
	}
	return 0, nil
}

func (f *FSMParser) getVarFloatVal(s string) (float32, error) {
	vIdx := strings.Index(s, "FLOAT")
	if vIdx > -1 {
		vIdx += 5
		cval := strings.Trim(s[vIdx:], " \t\n")
		if len(cval) == 0 {
			return float32(0), nil
		}
		rv, err := strconv.ParseFloat(cval, 32)
		if err != nil {
			return 0, err // Failed to parse
		}
		return float32(rv), nil
	}
	return 0, nil
}

func (f *FSMParser) getVarStringVal(s string) (string, error) {
	vIdx := strings.Index(s, "STRING")
	if vIdx > -1 {
		vIdx += 6
		cval := strings.Trim(s[vIdx:], " \t\n")
		return cval, nil
	}
	return "", nil

}

func (f *FSMParser) processVarLine(s string) {
	// Insure that there is a valid variable type
	ValidTypes := []string{"INT", "FLOAT", "STRING"}
	for _, v := range ValidTypes {
		if strings.Contains(s, v) {
			// We know which one this is!
			var tvar Var
			tvar.Name = f.getVarName(v, s)
			switch v {
			case "INT":
				// Process int
				tvar.VType = VInt
				tv, err := f.getVarIntVal(s)
				if err != nil {
					log.Fatal(err)
				}
				tvar.curval = tv

			case "FLOAT":
				// Process float
				tvar.VType = VFloat
				tv, err := f.getVarFloatVal(s)
				if err != nil {
					log.Fatal(err)
				}
				tvar.curval = tv
			case "STRING":
				// Process string
				tvar.VType = VString
				tv, err := f.getVarStringVal(s)
				if err != nil {
					log.Fatal(err)
				}
				tvar.curval = tv

			}
			f.Pmap.FSMVars = append(f.Pmap.FSMVars, tvar)
		}
	}
}

func (f *FSMParser) processVars(i int) (int, error) {
	for n, v := range f.FsmText[i:] {
		if f.checkTerminus(v) {
			return n + i, nil
		}
		f.processVarLine(v)
	}
	return i, nil
}

func (f *FSMParser) checkLabel(s string) bool {
	for _, v := range f.Pmap.FSMLabels {
		if v.Name == s {
			return true
		}
		sl := len(s) - 1
		if s[sl] == ':' {
			if v.Name == s[0:sl] {
				return true
			}
		}
	}
	// Is this an unknown label?
	return false
}

func (f *FSMParser) checkFunc(s string) bool {
	for _, v := range f.Pmap.FSMFuncs {
		if v.Name == s {
			return true
		}
	}
	return false
}

func (f *FSMParser) addFunc(s []string) {
	// Initially just add func into list
	// TODO: Parse line
	f.Pmap.FSMFuncs = append(f.Pmap.FSMFuncs, &FSMFunc{Name: s[0]})
}

func (f *FSMParser) processExecLine(s string) {
	// Check to see if we've got a function registered with this name. If it isn't, this must be a label
	// Before doing this, we'll need to parse the line into it's individual bits
	if len(s) == 0 {
		return
	}
	if s[0] == '#' {
		return
	}
	tln := strings.Fields(s)

	if len(tln[0]) == 0 {
		return
	}
	if f.checkLabel(tln[0]) {
		// This is a label
		// Add label address to jump table
		fmt.Printf("Adding label %s to jump table\n", tln[0])

	} else {
		// This is a function
		// Add function to jump table
		fmt.Printf("Adding function %s\n to jump table\n", tln[0])
		if !f.checkFunc(tln[0]) {
			f.addFunc(tln)
		}
	}
}

func (f *FSMParser) buildLabelsList(i int) {
	// Exec block is first, and **MUST** occur after Vars (if any) if exec is elsewhere, the states above won't be processed.
	// Yes, I can fix that... However, doing it this way, allows me to quickly comment out a chunk of the state machine without
	// actually commenting it out. This is in fact by design (hence why current line number is passed in)
	f.Pmap.FSMLabels = make(FSMLabelMap, 0)
	for _, v := range f.FsmText[i:] {
		l := strings.Trim(v, " \t\n")
		ll := len(l)
		if ll > 1 {
			ll--
			if l[ll] == ':' {
				ln := l[0:ll]
				if _, ok := f.Pmap.FSMLabels[ln]; !ok {
					f.Pmap.FSMLabels[l] = &FSMLabel{Name: ln}
				}
			}
		}

	}
}

func (f *FSMParser) processExecBlock(i int) (int, error) {
	f.buildLabelsList(i)
	for n, v := range f.FsmText[i:] {
		if f.checkTerminus(v) {
			return n + i, nil
		}
		f.processExecLine(v)
	}
	return i, nil
}

func (f *FSMParser) PrintPmap() {
	printVars := func() {
		VT := []string{"INT", "FLOAT", "STRING"}
		fmt.Printf("\tVars\n")
		for _, v := range f.Pmap.FSMVars {
			fmt.Printf("\t\t%s\t%s\t%v\n", v.Name, VT[int(v.VType)], v.curval)
		}
		fmt.Printf("----------------------\n")
	}
	printLabels := func() {
		fmt.Printf("\tLabels\n")
		for k, v := range f.Pmap.FSMLabels {
			fmt.Printf("\t\t%s\t%v\n", k, *v)
		}
		fmt.Printf("----------------------\n")
	}
	printFuncs := func() {
		fmt.Printf("\tFunctions\n")
		for _, v := range f.Pmap.FSMFuncs {
			fmt.Printf("\t\t%v\n", *v)
		}
		fmt.Printf("----------------------\n")
	}
	fmt.Printf("Pmap:\n")
	fmt.Printf("----------------------\n")
	printVars()
	printLabels()
	printFuncs()
}

func (f *FSMParser) BuildFSMElements() (bool, error) {
	// Iterate through state file, build variable list, label list, and function list
	var SkipLine int
	var err error
	for i, v := range f.FsmText {
		if len(v) == 0 {
			continue
		}
		if v[0] == '#' {
			continue
		}
		if i < SkipLine {
			continue
		}
		if strings.Contains(v, ":VARS") {
			SkipLine, err = f.processVars(i + 1)
			fmt.Printf("Variables are: %v\n", f.Pmap.FSMVars)
			if err != nil {
				return false, err
			}
			continue
		}
		if strings.Contains(v, ":EXEC") {
			SkipLine, err = f.processExecBlock(i + 1)
			if err != nil {
				return false, err
			}
			continue
		}
		f.processLine(i)
	}
	f.PrintPmap()
	return true, nil
}

func (f *FSMParser) processLine(l int) {
	inL := strings.Trim(f.FsmText[l], " \t\n")
	if f.checkReserved(inL) {
		return
	}
	if f.checkLabel(inL) {
		return
	}
	f.processExecLine(inL)

}

func (f *FSMParser) Parse() (bool, error) {
	if succ, err := f.BuildFSMElements(); (err == nil) && !succ {
		// Build jump table
		return succ, err
	}
	return true, nil
}
