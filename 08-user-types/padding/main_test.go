package padding

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test(t *testing.T) {

	fmt.Println("bool:", unsafe.Sizeof(true))  // 1
	fmt.Println("int:", unsafe.Sizeof(1))      // 8 на моей машине
	fmt.Println("string:", unsafe.Sizeof("A")) // 16 (длина + указатель)

	var x struct {
		a  bool   // 1
		b  bool   // 1
		c  string // 16
		bb int32  // 1
	}

	fmt.Println("x size:", unsafe.Sizeof(x)) // 24!
	fmt.Println("offsets:",
		unsafe.Offsetof(x.a), // 0
		unsafe.Offsetof(x.b), // 1
		unsafe.Offsetof(x.c)) // 8

}
