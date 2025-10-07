package fl

import "os"

type Flag struct {
	value       bool   
	description string 
}

var flags = make(map[string]*Flag)

func Parse() {
	for _, arg := range os.Args[1:] {
		if flagRef, ok := flags[arg]; ok {
			flagRef.value = true
		}
	}
}

func Bool(cmd string, value bool, description string) *bool {

	f := &Flag{
		value,
		description,
	}
	
	flags[cmd] = f

	return &f.value
}
