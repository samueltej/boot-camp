package fl

import "os"

type Flag struct {
	value       bool   
	description string 
}


var flags = make(map[string]*Flag)

func Parse() {

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if flagRef, exists := flags[arg]; exists {
			flagRef.value = true
		}
	}
}

func Bool(cmd string, value bool, description string) *bool {

	f := &Flag{
		value:       value,
		description: description,
	}

	flags[cmd] = f

	return &f.value
}
