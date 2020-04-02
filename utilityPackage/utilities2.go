package utilities

import (
	"fmt"
)

func PlayWithPackagesAndLocalFunctions(msg string) {
	fmt.Println("Hello from #2!", msg)

	// note: we can call a function declared in a different file of this package
	// even if it is not exported outside of the package
	sayHello3()
}