package foo

import "fmt"

func printFoo() string {
	s1 := fmt.Sprintf("s1: %s", "foo.func: PrintFoo()")
	return s1
}
