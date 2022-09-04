# Golang

### Contents

- [Terminal Stuff](#terminal-stuff)
- [Libaries](#libaries)
- [File Stuff](#file-stuff)
- [variable](#variable)
- [array](#array)
- [loops](#loops)
- [Maps](#maps)
- [Struct](#struct)
- [Controll flow](#controll-flow)
- [Switch statement](#switch-statement)
- [Short hand Switch statement](#short-hand-switch-statement)
- [Functions](#functions)
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
Wen we append to a slice we do like below
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
Due to Go being a pass by value language, modifying a normal array parameter won’t create permanent change. Sometimes this can be useful in performing local calculations.
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

- Create an empty map with `make`
```variableName := make(map[keyType]valueType)```

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
