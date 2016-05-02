package txb

import (
	"fmt"
)

func ExampleComputeChecksum() {
	sm := computeChecksum("vk_id", "app_id", "app_skey")

	// base 16, with lower-case letters for a-f
	str := fmt.Sprintf("%x", sm)

	fmt.Println(str)
	
	// Output:
	// 3830653031643662666633376464616666646165363532636565316466393062
}
