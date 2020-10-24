package anti_apple

import (
	"fmt"
	"runtime"
)

func init() {
	for runtime.GOOS == "darwin" {
		fmt.Println("I don't like Apple and its products, and I forbid using this code in any Apple product")
	}
}
