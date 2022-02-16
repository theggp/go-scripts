package main

import (
	"fmt"
	"math"
)



func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func TutorialMethods(){

  //Using a Methods
  v := Vertex{3, 4}
	fmt.Println(v.Abs())

  //Pointer receivers -> used to modify the object, not a copy of it
  // Also more efficient (don't need to create copies)
  v.Scale(10)
	fmt.Println(v.Abs())

  // If you want to do the same with a function you need the & operator
  Scale(&v, 10)
  fmt.Println(v.Abs())


  // Functions that take an arguments must take the correct type
  // While methods with value receivers can take both value or pointers to value
  // AbsFunc(v) OK, AbsFunct(&v) Not OK
  // v.Abs() OK, p := &v; p.Abs() still OK



  // For an explanation of interfaces see the doc
  // todo
  // Interfaces to detect underlying BasicTypesPlayground
  // todo

  // Stringers, a type that has the method String
  a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

  // Define a custom Type
  type IPAddr [4]byte

  // String exercise
  // s := make([]string, len(ip))
  // for i, val := range ip {
  //   s[i] = strconv.Itoa(int(val))
  // }
  // return strings.Join(s, ".")

  // byte to int int(val), byte to string string(val)

  // Todo Errors
  // Todo Readers


}
