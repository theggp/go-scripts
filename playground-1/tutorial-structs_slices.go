package main

import (
	"fmt"
	"strings"
)


type Vertex struct {
	X float64
	Y float64
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func TutorialStruct() {

	defer fmt.Println("This concludes the tutorial")

	// Pointers
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j


	// Struct, a collection of fields
	my_struct := Vertex{1, 2}
	my_struct.X = 3
	struct_p := &my_struct
	struct_p.Y = 4		// pointers of strucsts don't need deference
	fmt.Println(my_struct)
	my_struct2 := Vertex{X: 2}	// only one variable is initialized
	fmt.Println(my_struct2)


	// Arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(primes)

	// Slices, a slice does not store data, it just describe a section of an array
	var s []int = primes[1:4]
	fmt.Println(s)			// len(s), cap(s)

	// Slice literals, is like an array lioteral without the lenght
	q := []int{2, 3, 5, 7, 11, 13} 	// you can also create arrays/slices of structs
	fmt.Println(q)
	fmt.Println(q[:3])	// When slicing you can omit the low or high bounds
	fmt.Println(q[3:])
	fmt.Printf("len=%d cap=%d %v\n", len(q[:3]), cap(q[:3]), q[:3])

	var slice_1 []int
	fmt.Println(slice_1, len(slice_1), cap(slice_1))
	if slice_1 == nil {
		fmt.Println("The slice is nil!")
	}

	a1 := make([]int, 5)
	printSlice("a", a1)
	b1 := make([]int, 0, 5)
	printSlice("b", b1)
	c1 := b1[:2]
	printSlice("c", c1)
	d1 := c1[2:5]
	printSlice("d", d1)

	// Slices of Slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to a Slice
	var s1 []int
	s1 = append(s1, 0)
	s1 = append(s1, 1, 2, 3)
	printSlice("Appending", s1)


	// Example of image creation
	dx, dy := 255, 255
	image := make([][]uint8, dy)
	for i:=0; i<dy; i++ {
		line := make([]uint8, dx)
		for j:=0; j<dx; j++{
			//line[j] = (uint8(i)+uint8(j))/2
			line[j] = (uint8(j)^uint8(i))
		}
		image[i] = line
	}

	// Range
	var pow = []int{1, 2, 4, 8, 16}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Maps
	// var m map[string]Vertex
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])


	// Maps literals
	var m2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},	// don't need to specify this is a Vertex
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(m2)

	// Mutating Maps
	m3 := make(map[string]int)
	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])
	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])
	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])
	v, ok := m3["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// Closures here: https://go.dev/tour/moretypes/25




}

// Functions can be passed as arguments, same as Values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counter := make(map[string]int)
	for _, v := range words{
		if val, ok :=  counter[v]; ok {
			counter[v] = val+1
			} else {
			counter[v] = 1
			}
		}
	return counter
}
