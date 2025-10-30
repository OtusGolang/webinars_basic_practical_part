package padding

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPaddingOffsets(t *testing.T) {

	fmt.Println("bool:  ", unsafe.Sizeof(true)) // 1
	fmt.Println("int:   ", unsafe.Sizeof(1))    // 8 на моей машине
	fmt.Println("string:", unsafe.Sizeof("A"))  // 16 (длина + указатель)

	var x struct {
		a  bool   // 1
		b  bool   // 1
		c  string // 16
		bb int64  // 1
	}

	fmt.Println("x size:", unsafe.Sizeof(x)) // 24!
	fmt.Println("offsets:",
		unsafe.Offsetof(x.a), // 0
		unsafe.Offsetof(x.b), // 1
		unsafe.Offsetof(x.c)) // 8

}

func TestPaddingCases(t *testing.T) {
	var x1 struct {
		a int8  // 1
		b int16 // 2
	}
	fmt.Println("x1 size:", unsafe.Sizeof(x1)) // 4

	var x2 struct {
		a1 int8  // 1
		a2 int8  // 1
		b  int16 // 2
	}
	fmt.Println("x2 size:", unsafe.Sizeof(x2)) // 4

	var x3 struct {
		a int8  // 1
		b int32 // 4
	}
	fmt.Println("x3 size:", unsafe.Sizeof(x3)) // 8

	var x4 struct {
		a int8  // 1
		b int64 // 8
	}
	fmt.Println("x4 size:", unsafe.Sizeof(x4)) // 16

	var x5 struct {
		a int8 // 1
		// a  int16 // 2
		b1 bool // 1/8
		b2 bool // 1/8
		b3 bool // 1/8
		b4 bool // 1/8
		b5 bool // 1/8

	}
	fmt.Println("x5 size:", unsafe.Sizeof(x5)) // 2

	var x6 struct {
		b int16
		a [5]bool
	}
	fmt.Println("x6 size:", unsafe.Sizeof(x6)) // 8
}
