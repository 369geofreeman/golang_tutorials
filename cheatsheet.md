# Golang

### Contents

- [Hints](#hints)
- [Types](#types)
- [Terminal Stuff](#terminal-stuff)
- [Libaries](#libaries)
- [File Stuff](#file-stuff)
- [variable](#variable)
- [Stringers](#stringers)
- [array](#array)
- [loops](#loops)
- [Maps](#maps)
- [Struct](#struct)
- [Interface](#interface)
- [Controll flow](#controll-flow)
- [Switch statement](#switch-statement)
- [Short hand Switch statement](#short-hand-switch-statement)
- [Functions](#functions)
- [Variadic Functions](#variadic-functions)
- [Addresses](#addresses)
- [Pointers](#pointers)
- [Dereferencing](#dereferencing)
- [concurrency Go routines](#concurrency-go-routines)
- [Example of a file](#example-of-a-file)

<br>
<br>
<br>
<br>
<br>
<hr>


## Hints

- Get an objects type | fmt.Printf("%T\n", values)
- create error | errors.New(string)

- Get the current time | time.Now()
- Get datetime from string | myDate, err := time.Parse("2006-01-02 15:04:01", "July 25, 2022 13:45:00")
- get hour from time.time | now.Hour()

- convert float64 to float32 | f64 := float64(f32)
- Float tp precision | "%.2f"
- Use import "strconv" to covert between primative types | var s string = strconv.Itoa(number)

- String to lower/upper | strings.ToLower("Gopher")
- Strip all white space from a sting | strings.ReplaceAll(randomString, " ", "")

- Square ints | math.Pow(num, 2)  // returns float64
- Square floats to int | int(math.Pow(num, 2))

- Generate random number | n := rand.Intn(100) // n is a random int, 0 <= n < 100
- Generate random number between two ints | rand.Intn(max - min) + min
- Generate random float | f := rand.Float64() // f is a random float64, 0.0 <= f < 1.0
- Generate a random float between two floats | r := min + rand.Float64() * (max - min)
- Setting a seed for random numers | rand.Seed(time.Now().UnixNano())
- shuffle a slice 
```
x := []string{"a", "b", "c", "d", "e"}
// shuffling the slice put its elements into a random order
rand.Shuffle(len(x), func(i, j int) {
	x[i], x[j] = x[j], x[i]
})
```


<br>
<br>
<br>
<br>
<br>
<hr>

## Types

- Ints

| type   |                                                                      |
|--------|----------------------------------------------------------------------|
| uint8  | Unsigned 8-bit integers (0 to 255)                                   |
| uint16 | Unsigned 16-bit integers (0 to 65535)                                |
| uint32 | Unsigned 32-bit integers (0 to 4294967295)                           |
| uint64 | Unsigned 64-bit integers (0 to 18446744073709551615)                 |
| int8   | Signed 8-bit integers (-128 to 127)                                  |
| int16  | Signed 16-bit integers (-32768 to 32767)                             |
| int32  | Signed 32-bit integers (-2147483648 to 2147483647)                   |
| int64  | Signed 64-bit integers (-9223372036854775808 to 9223372036854775807) |


- Floats

| type       |                                                                      |
|------------|----------------------------------------------------------------------|
| float32    | IEEE-754 32-bit floating-point numbers                               |
| float64    | IEEE-754 64-bit floating-point numbers                               |
| complex64  | Complex numbers with float32 real and imaginary parts                |
| complex128 | Complex numbers with float64 real and imaginary parts                |


- Other types

| type    |                                                                        |
|---------|------------------------------------------------------------------------|
| byte    | same as uint8                                                          |
| rune    | same as int32                                                          |
| uint    | 32 or 64 bits                                                          |
| int     | same size as uint                                                      |
| uintptr | an unsigned integer to store the uninterpreted bits of a pointer value |



- Zero values

| Type      | Zero Value |
|-----------|------------|
| boolean   | false      |
| numeric   | 0          |
| string    | ""         |
| pointer   | nil        |
| function  | nil        |
| interface | nil        |
| slice     | nil        |
| channel   | nil        |
| map       | nil        |

**Runes**

The rune type in Go is an alias for int32. Given this underlying int32 type, the rune type holds a signed 32-bit integer value. However, unlike an int32 type, the integer value stored in a rune type represents a single Unicode character.

Variables of type rune are declared by placing a character inside single quotes:
```myRune := '¿'```

- Since rune is just an alias for int32, printing a rune's type will yield int32:
```
myRune := '¿'
fmt.Printf("myRune type: %T\n", myRune)
// Output: myRune type: int32
```

- Similarly, printing a rune's value will yield its integer (decimal) value:
```
myRune := '¿'
fmt.Printf("myRune value: %v\n", myRune)
// Output: myRune value: 191
```

- To print the Unicode character represented by the rune, use the %c formatting verb:
```
myRune := '¿'
fmt.Printf("myRune Unicode character: %c\n", myRune)
// Output: myRune Unicode character: ¿
```

- To print the Unicode code point represented by the rune, use the %U formatting verb:
```
myRune := '¿'
fmt.Printf("myRune Unicode code point: %U\n", myRune)
// Output: myRune Unicode code point: U+00BF
```

- Runes and Strings

Strings in Go are encoded using UTF-8 which means they contain Unicode characters. Since the rune type represents a Unicode character, a string in Go is often referred to as a sequence of runes. However, runes are stored as 1, 2, 3, or 4 bytes depending on the character. Due to this, strings are really just a sequence of bytes. In Go, slices are used to represent sequences and these slices can be iterated over using range.

Even though a string is just a slice of bytes, the range keyword iterates over a string's runes, not its bytes. In this example, the index variable represents the starting index of the current rune's byte sequence and the char variable represents the current rune:

```
myString := "❗hello"
for index, char := range myString {
  fmt.Printf("Index: %d\tCharacter: %c\t\tCode Point: %U\n", index, char, char)
}
// Output:
// Index: 0	Character: ❗		Code Point: U+2757
// Index: 3	Character: h		Code Point: U+0068
// Index: 4	Character: e		Code Point: U+0065
// Index: 5	Character: l		Code Point: U+006C
// Index: 6	Character: l		Code Point: U+006C
// Index: 7	Character: o		Code Point: U+006F
```

Since runes can be stored as 1, 2, 3, or 4 bytes, the length of a string may not always equal the number of characters in the string. Use the builtin len function to get the length of a string in bytes and the utf8.RuneCountInString function to get the number of runes in a string:
```
import "unicode/utf8"

myString := "❗hello"
stringLength := len(myString)
numberOfRunes := utf8.RuneCountInString(myString)

fmt.Printf("myString - Length: %d - Runes: %d\n", stringLength, numberOfRunes)
// Output: myString - Length: 8 - Runes: 6
```

<br>
<br>
<br>
<br>
<br>
<hr>

## Terminal Stuff

- compile code. in terminal run => **go build nameOfFile.go**

- use **ls** to make sure the file is there

- toexecute the comppiled file run =>  **./nameOfFile**

- run file without executing, run => **run nameOfFile.go**

- find information about a package, run => **go doc packageName**

<br>
<br>
<br>
<br>
<br>
<hr>

## Libaries

- math
- math/rand		<-- library that helps us generate a random integer => rand.Intn(100)
- time			<--- time.Now() for current time
- fmt			<-- format print [.Print(), .Println(), .Sprint(), .Scan()]


**fmt**

- .Print() => no extra formating 
- .Println() => extra formating, spaces new lines etc 
- .Sprint(), => does now print but adds variables. Returns the value instead of preints
- .Scan() => gets user input `fmt.Scan(&input)`

**random**

To use random we need to set a seed for it to run, an example of this could be using the current time at the start of the file or function like so:
```
package main

import (
	"fmt"
  "math/rand"
  "time"
)

func main() {
	// Set random seed here
  rand.Seed(time.Now().UnixNano())
  
  amountLeft := rand.Intn(10000)
  
  fmt.Println("amountLeft is: ", amountLeft)
  
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
  } else {
    fmt.Println("Where did all my money go?")
  }
}
```
 
<br>
<br>
<br>
<br>
<br>
<hr>

## File Stuff 

Always start the main file with
```
package main
```

import libraies like
```
import "fmt"
```
or multiple libraies like
```
import (
    "fmt",
    "time",
    m "math"    <- this means we imported math as m
)
```

<br>
<br>
<br>
<br>
<br>
<hr>

## **variable**

- var name string = "Jelly"					<-- standard way
- name := "jelly"							<-- go automatically infers the type
- name, age, likeGo := "Jelly", 29, true	<-- multiple variables on the same line 

<br>
<br>
<br>
<br>
<br>
<hr>

## Stringers

- Stringer is an interface for defining the string format of values.

The interface consists of a single String method:
```
type Stringer interface {
    String() string
}
```

Types that want to implement this interface must have a String() method that returns a human-friendly string representation of the type. The fmt package (and many others) will look for this method to format and print values.

**Example: Distances**

Assume we are working on an application that deals with geographical distances measured in different units. We have defined types `DistanceUnit` and `Distance` as follows:
```
type DistanceUnit int

const (
	Kilometer    DistanceUnit = 0
	Mile         DistanceUnit = 1
)
 
type Distance struct {
	number float64
	unit   DistanceUnit
} 
```
In the example above, Kilometer and Mile are constants of type `DistanceUnit`.

These types do not implement interface Stringer as they lack the String method. Hence fmt functions will print Distance values using Go's "default format":
```
mileUnit := Mile
fmt.Sprint(mileUnit)
// => 1
// The result is '1' because that is the underlying value of the 'Mile' constant (see constant declarations above) 

dist := Distance{number: 790.7, unit: Kilometer}
fmt.Sprint(dist)
// => {790.7 0}
// not a very useful output!
```

In order to make the output more informative, we implement interface Stringer for `DistanceUnit` and Distance types by adding a String method to each type:
```
func (sc DistanceUnit) String() string {
	units := []string{"km", "mi"}
	return units[sc]
}

func (d Distance) String() string {
	return fmt.Sprintf("%v %v", d.number, d.unit)
}
```

fmt package functions will call these methods when formatting Distance values:
```
kmUnit := Kilometer
kmUnit.String()
// => km

mileUnit := Mile
mileUnit.String()
// => mi

dist := Distance{
	number: 790.7,
	unit: Kilometer,
}
dist.String()
// => 790.7 km
```

<br>
<br>
<br>
<br>
<br>
<hr>

## **array**

An array is a collection of data elements of the same type, where we can access each element by an index.
- Storing many pieces of input
- Storing related collections of values
- Performing mathematical operations on lists of numbers

When we declare a variable in Go, the compiler:

- Finds space in memory for that variable
- Associates the variable with a name

We can have the compiler determine the length automatically using ... ellipsis syntax.
```triangleAngles := [...]int{30, 60, 90}```

- creating an array
```
menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}
```

- Getting element at index from array
```arr[0]```

- Initialising and adding to an array
```
myArray = [3]string
myArray[0] = "a"
myArray[1] = "b"
myArray[2] = "c"
```

- Slicing arrays
Slices are a data collection type similar to arrays, but they have the ability to change their size
```
// Each of the following creates an empty slice
var numberSlice []int
stringSlice := []string{}
 
// The following creates a slice with elements
names := []string{"Kathryn", "Martin", "Sasha", "Steven"}
```

- Capacity and length

length is the amount an array or slice is holding, capacity is the amount oif space allocated to that the slice could hold before needing to resize itself.

```
slice := []string{"Fido", "Fifi", "FruFru"}
// The slice begins at length 3 and capacity 3

fmt.Println(slice, len(slice), cap(slice))
// [Fido Fifi FruFru] 3 3

slice = append(slice, "FroFro")
// After appending an element when the slice is at capacity
// The slice will double in capacity, but increase its length by 1

fmt.Println(slice, len(slice), cap(slice))
// [Fido Fifi FruFru FroFro] 4 6
```

- Append
When we append to a slice we do like below
```
func main() {
    myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
    myTutors = append(myTutors, "Josh")
}
```

- Arrays in functions

We can pass arrays or slices into functions as parameters. To pass an array parameter into a function, we provide a local name, square brackets, and the data type. The difference between slice and array parameters is whether the number of elements is stated:

```
// Array
func printFirstLastArray(array [4]int) {
    fmt.Println("First", array[0])
    fmt.Println("Last", array[3])
}
 
// Slice
func printFirstLastSlice(slice []int) {
    length := len(slice)
    if (length > 0) {
        fmt.Println("First", slice[0])
        fmt.Println("Last", slice[length-1])
    }
}
```
Due to Go being a pass-by-value language, modifying a normal array parameter won’t create permanent change. Sometimes this can be useful in performing local calculations.
```
// Changes to the array will only be local to the function
func changeFirst(array [4]int, value int) {
    array[0] = value
}
```


<br>
<br>
<br>
<br>
<br>
<br>

## **loops**

- Classic loop example
```
for x := 0; x < 10; x++ {
	fmt.Println(x)
}   
```

- For as a While Loop (while loop - loop until a condition is met)
```
number := 0 // Initialize a variable to be used inside the loop
for number < 5 {
  fmt.Println(number)
  number++ // Update the variable being used
}
```

- Looping with arrays. we can use teh `range` keyword to loop arrays
```
letters := []string{"A", "B", "C", "D"}
for index, value := range letters {
  fmt.Println("Index:", index, "Value:", value)
}
```

<br>
<br>
<br>
<br>
<br>
<hr>


## **Maps**

- To create a map, you can do:
```
  // With map literal
  foo := map[string]int{}
```
or
```
  // or with make function
  foo := make(map[string]int)
```
- Here are some operations that you can do with a map
```
  // Add a value in a map with the `=` operator:
  foo["bar"] = 42
  // Here we update the element of `bar`
  foo["bar"] = 73
  // To retrieve a map value, you can use
  baz := foo["bar"]
  // To delete an item from a map, you can use
  delete(foo, "bar")
```


- Create map with values
```
variableName := map[keyType]valueType{
    name1: value1,
    name2: value2,
    name3: value3,
}
```
 
- Creating an empty map and adding values to it
```
func main() {
	myMap := map[string]int{}
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	fmt.Println(myMap)
}
```

- Accessing elements in a map 
```
count,status := inventory["sporks"]
 
if status {
  fmt.Println("we have %d sporks!", count)
} else {
  fmt.Println("what is a spork?")
}
```

- Adding or updating

```customers["George"] = 10.50```


- Deleting from a map

```delete(yourMap, keyValueToDelete)```

- check key in map
```
if val, ok := dict["foo"]; ok {
    //do something here
}
```

- delete key value from map
```delete(map, key)```



<br>
<br>
<br>
<br>
<br>
<hr>

## Struct

We often have many variables related to each other, but managing using them all together can be a headache. Luckily, Go provides us with a way to group several variables into one custom data type. These types make the code cleaner, more intuitive, and less error-prone.

In Go, grouping together related variables is done using a struct

- Basic example
```
type Country struct {
  name string
  capital string
  latitude float32
  longitude float32
}
```
- Creating an Instance of a Struct

```p1 := Point{x: 10, y: 12}```
if we omit the fields GO will initialise the values to 0, '', false etc
```
p1 := Point{x: 10}
// y will be set to 0

p1 := Point{}
// x and y will be set to 0
```

- Accessing and Modifying Struct Variables

Let’s say we have an instance of the Student struct:
```john := Student{"John", "Smith", 14, 9}```
We can access individual fields within struct using the name of the variable, a ., and the name of the field. We could access John’s first name like so:
```fmt.Println(john.firstName)```
We can change the value of a field with an assignment statement:
```john.age = 15```


- Functions that Access a Struct

We can use functions to capture logic involving our structs and simplify it.

Structs will often have important operations that can be performed on them. For example, with a struct representing a geometric shape, it would be natural to have functions that compute its area and perimeter.

Let’s say we have a struct describing a rectangle. The rectangle struct will contain two fields: the length and the width. We define this struc

```
type Rectangle struct {
  length float32
  width  float32
}
```
We can define a function that computes the area of the rectangle; the product of the length and the width.

```
func (rectangle Rectangle) area() float32 {
  return rectangle.length * rectangle.width
}
```
The key thing to notice is the line (rectangle Rectangle). This line signals to Go that the area() function belongs to the Rectangle struct. Note that functions associated with a struct are written outside of the struct!

If we have an instance of Rectangle called rect, we can call the area() function like so:
```
rect.area()
```
Defining a function in this way will only pass in a copy of the rectangle: that is, we will not be able to use the function to alter the value of a field!

If we want to write a function that allows us to modify the value of a struct field, we have to pass in a pointer to a struct


- Pointers to a Struct

Without pointers, when a variable is passed into a function, only a copy of it is used inside the function. We can use pointers to modify values in our structs within a function. But how do we get a pointer to our struct?

Let’s explore this concept using the following struct as an example:
```
type Employee struct {
  firstName string
  lastName string
  age int
  title string
}
```
We must first create an instance of Employee and then we create a pointer that will point to this instance. This is done like so:
```
func main() {
  steve := Employee{“Steve”, “Stevens”, 34, “Junior Manager”}
  pointerToSteve := &steve
}
```
We can now use this pointer to change the values of the fields for steve. There are two ways to do this in Go:
```
(*pointerToSteve).firstName
```
Or a simpler, recommended method:
```pointerToSteve.firstName ```

We can use these pointers to modify structs in our functions. Consider the following example:
```
func (rectangle *Rectangle) modify(newLength float32){
  rectangle.length = newLength
}
```
Notice that just inside the function modify() that rectangle is also a pointer. It is dereferenced without the use of the dereferencing operator just like pointerToSteve!

We now have the ability to change Struct values in our functions!

- **FULL POINTER FOR STRUCT EXAMPLE
```
package main

import "fmt"

type Triangle struct {
	height float32
	base float32
}

// Checkpoint 1 code goes here
func (triangle *Triangle) updateBase(newVal float32) {
  triangle.base = newVal
}

func main() {

  triangle := Triangle{10, 4}
  
  // Call the function here
  triangle.updateBase(13)

  fmt.Println(triangle)

}
```


- Arrays of structs

What can we do when dealing with many structs of the same type? We can use them in an array together! Let’s say we want to create an array of the following points: 
```{1, 1} {7, 27} {12, 7} {9, 25}```

We create an array of Points like so:
```
points := []Point{{1, 1}, {7, 27}, {12, 7}, {9, 25}}
```
If the points have names, we can also create the array like this:
```
a = {1, 1}
b = {7, 27}
c = {12, 7}
d = {9, 25}

points := []Point{a, b, c, d}
```
We can access the contents of this array like we would for an ordinary array. We can also access and modify the fields of each of the array elements.
```
points := []Point{{1, 1}, {7, 27}, {12, 7}, {9, 25}}
 
fmt.Println(points[0]) // Output will be {1, 1}
points := []Point{{1, 1}, {7, 27}, {12, 7}, {9, 25}}
 
points[1].x = 8
points[1].y = 16
 
fmt.Println(points[1]) // Output will be {8, 16}
```
Arrays of structs allow us access many instances of our structs together in our programs!


- Nested Structs

When we have complex groups of fields in our structs, they can be combined into their own struct! For example, in the Employee struct, we have two separate fields for the first and last name of the employee. We can combine those two strings into their own struct called Name:

```
type Name struct{
  firstName string
  lastName string
}
 
type Employee struct{
  name Name
  age int
  title string
}
```
We create an instance of Employee like so:
```
carl := Employee{Name{“Carl”, “Carlson”}, 32, “Engineer”}
```
To access the fields of the nested struct (Name in this case), we chain together the field accesses like so:
```
fmt.Println(carl.name.lastName) // Output will be “Carlson”
```
We can also define the employee struct with the Name struct anonymously like so:
```
type Employee struct{
  Name
  age int
  title string
}
```
Notice how the Name file has no associated variable name with it.

Composing a struct in this way allows us to access the firstName and lastName fields directly from the Employee struct.
```
carl := Employee{Name{“Carl”, “Carlson”}, 32, “Engineer”}
 
fmt.Println(carl.firstName) // Output will be “Carl”
fmt.Println(carl.lastName) // Output will be “Carlson”
```
We, of course, cannot have two anonymous fields of the same type (i.e., two Name fields) as that would make it impossible to tell which field is being accessed (which firstName if two anonymous Name fields).

An anonymous field is used to field access easier and leads to cleaner code.

We now have the ability to organize struct information inside of existing structs!


<br>
<br>
<br>
<br>
<br>
<hr>

## Interface

An interface type is effectively a set of method signatures. Any type that defines those methods "implements" the interface implicitly. There is no implements keyword in Go. Here is what an interface definition might look like:
```
type InterfaceName interface {
    MethodOne() MethodOneReturnType
    MethodTwo(paramOne ParamOneType, paramTwo ParamTwoType) MethodTwoReturnType 
    ...
}
```
For example, here is the built-in error interface:
```
type error interface {
    Error() string
}
```
This means that any type which implements an `Error()` method which returns a string implements the error interface. This allows a function with return type error to return values of different types as long as all of them satisfy the error interface.

There is one very special interface type in Go: the empty interface type that contains zero methods. The empty interface type is written like this: `interface{}`. 
Since it has no methods, every type implements the empty interface type. This is helpful for defining a function that can generically accept any value. In that case, the function parameter uses the empty interface type.

<br>
<br>
<br>
<br>
<br>
<hr>


## **Controll flow**

Short hand if statment (controll flow)
```
if success:= true; success {
    fmt.Println("We're rich!")
} else {
    fmt.Println("Where did we go wrong?")
}
```

<br>
<br>
<br>
<br>
<br>
<hr>


## Switch statement
```
func main() 

	amountStolen := 50000
	numOfThieves := 5

	switch numOfThieves {
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
	  fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
	  fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}
}
```

## Short hand Switch statement
```
amountStolen := 50000

switch numOfThieves := 5; numOfThieves {
case 1:
    fmt.Println("I'll take all $", amountStolen)
default:
    fmt.Println("There's not enough to go around...")
```


<br>
<br>
<br>
<br>
<br>
<hr>


### Functions

- ```main()``` calls all function and runs all code (think python __main__ == "__name__")
- Can't return a function with declaring a return type
- Function parameters are like typing hints to show what the variables passsed in are


```
func computeMarsYears(earthYears int32) int {
  
  earthDays := earthYears * 365
  marsYears := earthDays / 687
  return marsYears
}
```

- We can combine similar parameters i fthey are of the same type

```func computeMarsYears(earthYears, earyhDays int32) int {}```

- We can return multiple things by declaring the return values in parenthesis
```
func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
  averageGrade := (midtermGrade + finalGrade) / 2
  var gradeLetter string
  if averageGrade > 90 {
    gradeLetter = "A"
  } else {
    gradeLetter = "F"
  }
 
  return gradeLetter, averageGrade 
}
```

- **defer** We can delay a function call to the end of the current scope by using the defer keyword

```
func calculateTaxes(revenue, deductions, credits float64) float64 {
  defer fmt.Println("Taxes Calculated!")
  taxRate := .06143
  fmt.Println("Calculating Taxes")
 
  if deductions == 0 || credits == 0 {
    return revenue * taxRate
  } else {
    return 0
  }
}
```
returns
```
Calculating Taxes
Taxes Calculated!
```


<br>
<br>
<br>
<br>
<br>
<hr>

### Variadic Functions

Usually, functions in Go accept only a fixed number of arguments. However, it is also possible to write variadic functions in Go.

A variadic function is a function that accepts a variable number of arguments.

If the type of the last parameter in a function definition is prefixed by ellipsis ..., then the function can accept any number of arguments for that parameter.

The variadic parameter must be the last parameter of the function.

Here is an example of an implementation of a variadic function.
```
func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)

    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            return
        }
    }

    fmt.Println(num, "not found in ", nums)
}

func main() {
    find(89, 90, 91, 95)
    // =>
    // type of nums is []int
    // 89 not found in  [90 91 95]

    find(45, 56, 67, 45, 90, 109)
    // =>
    // type of nums is []int
    // 45 found at index 2 in [56 67 45 90 109]

    find(87)
    // =>
    // type of nums is []int
    // 87 not found in  []
}
```


<br>
<br>
<br>
<br>
<br>
<hr>

### Pointers

## **Addresses**

To find a variable’s address we use the & operator followed by the variable name, like so:
```
x := "My very first address"
fmt.Println(&x) // Prints 0x414020
```
The 0x prefix is formatted in hexadecimal, which is a way of representing 16 digit numbers


## **Pointers**

No pointer arithmetic as this leads to dangerous and unpredictable behaviour

- Pointers are variables that specifically store addresses.

- We even set the data type of the addresses’ value for the pointer. For instance:
```
var pointerForInt *int
```
In the example above pointerForInt will store the address of a variable that has an int data type. To break it down further, the * operator signifies that this variable will store an address and the int portion means that the address contains an integer value.


- With pointerForInt initialized, we can assign it value like so:
```
var pointerForInt *int
minutes := 525600
pointerForInt = &minutes
 
fmt.Println(pointerForInt) // Prints 0xc000018038   
```

- We can also declare a pointer implicitly like other variables:
```
minutes := 55
pointerForInt := &minutes
```

**Pointer with string full example**
```
package main

import "fmt"

func main() {
	star := "Polaris"

	var starAddress *string = &star
 	fmt.Printf("The address of star is %v", starAddress)
}
```

## **Dereferencing**

- We can actually use our pointer to access the address and change its value! This action is called dereferencing or indirecting.
- We’ll need to use the * operator again on a pointer and then assign a new value like so:
```
lyrics := "Moments so dear" 
pointerForStr := &lyrics
 
*pointerForStr = "Journeys to plan" 
 
fmt.Println(lyrics) // Prints: Journeys to plan
```

In our example, we have our variables: lyrics that has the value of "Moments so dear" and pointerForStr which is a pointer for lyrics. Then we use the * operator on pointerForStr to dereference it and assign a new value of "Journeys to plan". When we check the value of lyrics it’s now "Journeys to plan"!

**Changing Values in Different Scopes**

If we want to change the value of x using a function, we’re going to need to first change our function:
```
func addHundred (numPtr *int) {
  *numPtr += 100
}
```
Our new function now has a parameter of a pointer for an integer. By passing the value of a pointer (which is an address) to addHundred(), we can also dereference the address and add 100 to its value. But now that addHundred() expects a pointer for an argument, we’re also going to need to change our main()! The complete code is as follows:
```
func addHundred (numPtr *int) {
  *numPtr += 100
}
 
func main() {
  x := 1
  addHundred(&x)
  fmt.Println(x) // Prints 101
}
```
Since addHundred() expects a pointer (and pointers are variables that store an address) the final touch was to provide addHundred() the address of x. With that, x is now 101!


<br>
<br>
<br>
<br>
<br>
<hr>


## **concurrency Go routines**

GO routines are function that can run at the same time as other functions by utilising multiple threads on a CPU

goroutines are better described as green threads, which dynamically map function execution to operating system threads as needed. goroutines might be executed asynchronously on the same operating system thread, which is especially useful in the case that they perform io operations

We use the `go` expression before calling a function to make it a go routine

- Example of a go routine calling the same function twice at the same time
```
package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	count("fish")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
```

- We can create many simultanious go routines as Go does a very good job with concurrency, but it won't make the program any faster as we are constrained by how many calls our CPU has.


<hl>

	
<br>
<br>
<br>
<br>
<br>
<hr>


## Example of a file

```
package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Printf("There is %v fuel remaining\n", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int

  switch {
    case planet == "Venus":
      fuel = 300000
    case planet == "Mercury":
      fuel = 500000
    case planet == "Mars":
      fuel = 700000
    default:
    fuel = 0
  }

  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Printf("Welcome to planet %v\n", planet)
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int{
  var fuelRemaining, fuelCost int

  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)

  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }

  return fuelRemaining
}

func main() {
  // Test your functions!

  // Create `planetChoice` and `fuel`
  var fuel int
  fuel = 1000000
  planetChoice := "Venus"

  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
  
  // And then liftoff!
  
}
```
