package variables

import (
	"fmt"
	"os"
)

func runConstAndEnvVariable() {
	const c = 300000
	fmt.Println("Speed of light is ", c, " km/s")

	// c = c * 3600
	// fmt.Println("Speed of light is ", c, " km/hr")
	for _, env := range os.Environ() {
		fmt.Println("Environment Variable is ", env)
	}
}
