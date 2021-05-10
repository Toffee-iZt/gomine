package common

import "fmt"

// Try performs the function with possible
// panic without stopping the program
func Try(function func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("caught: %v", r)
		}
	}()

	function()
	return
}
