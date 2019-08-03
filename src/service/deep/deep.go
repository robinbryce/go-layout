package deep

import "fmt"

func HelloDeep() {
	fmt.Println("Hello Deep Package")
}

// Run contains the startup code for the 'deep' example service
func Run() int {
	HelloDeep()
	return 0
}
