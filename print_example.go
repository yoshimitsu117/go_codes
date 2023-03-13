package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Three types of data in golang
	/**
	1.Numbers:
	---------------------------------------------------------------------------
	i Integers:
	-----------------------------
	int8	8-bit signed integer
	int16	16-bit signed integer
	int32	32-bit signed integer
	int64	64-bit signed integer
	uint8	8-bit unsigned integer
	uint16	16-bit unsigned integer
	uint32	32-bit unsigned integer
	uint64	64-bit unsigned integer
	int	Both int and uint contain same size, either 32 or 64 bit.
	uint	Both int and uint contain same size, either 32 or 64 bit.
	rune	It is a synonym of int32 and also represent Unicode code points.
	byte	It is a synonym of uint8.
	uintptr	It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.

	ii. Floating Point Numbers
	---------------------------
	float 32 bit
	float 64 bit

	iii. Complex Numbers
	---------------------------
	c1 := complex(10, 11)
	c1 := 10 + 11i

	2.Boolean
	--------------------------------------------------------------------------------
	flag bool

	3. String
	-------------------------------------------------------------------------------

	**/
	var flag bool
	var complexNum complex64 = 5 + 11i
	if reflect.TypeOf(complexNum).Name() == "complex64" {
		fmt.Println("The type of data is complex 64")
	}
	if reflect.TypeOf(complexNum) == nil {
		fmt.Println("No data type found")
	} else {
		fmt.Printf("%v", reflect.TypeOf(complexNum))
	}

	fmt.Printf("%v", complexNum)
	fmt.Print(flag)
	if flag == false {
		fmt.Printf("The answer is %v", flag)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%s", "geeksforgeeks")
	}
	fmt.Println("Hello There!")
}
