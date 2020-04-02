package utilities

import (
	"fmt"
)

// note that this is only available to other files within this package
func sayHello3() {
	fmt.Println("Hello from SayHello3 which is local to the utilities package")
}