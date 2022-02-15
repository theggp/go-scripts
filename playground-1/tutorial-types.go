package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"unsafe"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const(
  a_const int = 12
)

func BasicTypesPlayground() {

	fmt.Println("Some variable and a constant")
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
  fmt.Printf("+Type: %T Value: %v\n", a_const, a_const)

	a_int8 := int8(1<<7 - 1)
	a_uint8 := uint8(1<<8 - 1)
	a_int16 := int16(1<<15 - 1)
	a_uint16 := uint16(1<<16 - 1)
	a_int32 := int32(1<<31 - 1)
	a_uint32 := uint32(1<<32 - 1)
	a_int64 := int64(1<<63 - 1)
	a_uint64 := uint64(1<<64 - 1)
	a_int := int(1<<63 - 1)
	a_uint := uint(1<<64 - 1)

	// byte = alias for unit8
	// rune = alias for a_int32

	a_float32 := float32(1<<32+10100000000000000)
	a_float64 := float64(1<<64 - 1)
	// complex64 for complex numbers
	// complex128 for complex numbers

	fmt.Println("")
	fmt.Println("Types and max sizes")
	fmt.Printf("Type: %T,   Bytes: %d, Max Val: %v\n", a_int8, unsafe.Sizeof(a_int8), a_int8)
	fmt.Printf("Type: %T,  Bytes: %d, Max Val: %v\n", a_uint8, unsafe.Sizeof(a_uint8), a_uint8)
	fmt.Printf("Type: %T,  Bytes: %d, Max Val: %v\n", a_int16, unsafe.Sizeof(a_int16), a_int16)
	fmt.Printf("Type: %T, Bytes: %d, Max Val: %v\n", a_uint16, unsafe.Sizeof(a_uint16), a_uint16)
	fmt.Printf("Type: %T,  Bytes: %d, Max Val: %v\n", a_int32, unsafe.Sizeof(a_int32), a_int32)
	fmt.Printf("Type: %T, Bytes: %d, Max Val: %v\n", a_uint32, unsafe.Sizeof(a_uint32), a_uint32)
	fmt.Printf("Type: %T,  Bytes: %d, Max Val: %v\n", a_int64, unsafe.Sizeof(a_int64), math.MaxInt64)
	fmt.Printf("Type: %T, Bytes: %d, Max Val: %v\n", a_uint64, unsafe.Sizeof(a_uint64), a_uint64)
	fmt.Printf("Type: %T,    Bytes: %d, Max Val: %v\n", a_int, unsafe.Sizeof(a_int), a_int)
	fmt.Printf("Type: %T,   Bytes: %d, Max Val: %v\n", a_uint, unsafe.Sizeof(a_uint), a_uint)
	fmt.Printf("Type: %T,   Bytes: %d, Max Val: %v\n", a_float32, unsafe.Sizeof(a_float32), math.MaxFloat32)
	fmt.Printf("Type: %T,   Bytes: %d, Max Val: %v\n", a_float64, unsafe.Sizeof(a_float64), math.MaxFloat64)

	fmt.Println("")
	fmt.Println("Example of float rounding")
	x_float := float64(1)
	y_float := float64(1000000000000000000000000000000000000000000000000000000000)
	fmt.Println(x_float)
	x_float = x_float + y_float
	fmt.Println(x_float)
	x_float = x_float - y_float
	fmt.Println(x_float)

	fmt.Println("")
	fmt.Println("Example of inizialization values")
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("int:%v, float64: %v, bool: %v, string: %q\n", i, f, b, s)

}
