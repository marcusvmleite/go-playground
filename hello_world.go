package main

//Importing is file scoped
import (
	"fmt"
	somealias "golang-playground/test"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
	"unicode/utf8"
)

const ok = true

//File scope variables doesn't need to be used
var speed int

//All Go programs needs a main function
func main() {
	fmt.Println("Hello World!")
	fmt.Println(runtime.NumCPU())

	//_ is used for unused variables to make Go compiler happy
	var _ = "something"

	//Calling an exported function from test package
	//Here we are using the alias used to import, if
	//we hadn't created with alias then we should have called
	//like test.Print()
	somealias.Print()

	simple()
	variableDeclarations()
	converting()
	slices()
	arithmetic()
	strings()
	definedTypes()
	iotaExample()
	errorHandling()
	switchExample()
	loops()

	var dir, file string
	dir, file = path.Split("C:/Intel/Profiles/IntelGraphicsProfiles/test.exe")

	_dir, _file := path.Split("C:/Intel/Profiles/IntelGraphicsProfiles/test.exe")

	//We can ignore one of the values returned by Split function
	//using the underscore declaration '_'
	_, second := path.Split("C:/Intel/Profiles/IntelGraphicsProfiles/test.exe")

	fmt.Println(dir, file, _dir, _file, second)
}

//1. Go is STRONGLY TYPED!
//2. Every Go type has a zero value (e.g.: int has 0)
//3. Names should always start with letters or _
//4. Variables that start with uppercase letters are exported just like functions
func variableDeclarations() {
	var test int     //initialized to 0
	var _test int    //initialized to 0
	var heat float64 //initialized to 0
	var cond bool    //initialized to false
	var brand string //initialized to ""
	var now time.Time = time.Now()

	fmt.Println(test, _test, heat, cond, now)
	fmt.Printf("%q", brand)

	//Printing program arguments
	fmt.Printf("%#v\n", os.Args)

	//Another way to declare grouping. It's a good practice
	//to tell that they are somehow related
	var (
		name     string
		passport string
		age      int
	)

	fmt.Println(name, passport, age)

	//Inline
	var speed, velocity int
	fmt.Println(speed, velocity)

	//Declaration with initialization

	var init bool = true

	//We can infer
	var another = true

	//And even do this, that is called short declaration statement (:=)
	//You can't do short declaration in package-level! Only in function body!
	//ATTENTION: don't use short declaration when you don't
	//know which is the initial value, this is a bad practice.
	justOther := false
	//score := 0 DON'T
	//Instead
	var score int //as it will be initialized to 0

	//Another example: you want to declare 2 variables that you
	//know their initial values, don't do this
	//var height, weight = 10, 20
	height, weight := 10, 20 //Do this!

	fmt.Println(init, another, justOther, score, height, weight)

	//You can short declare multiple variables even
	//with different types
	one, two := 1, "2"

	fmt.Println(one, two)

	//Redeclaration
	var safe bool //Initialized do false
	//At least one of the variables needs to be a new one,
	//otherwise redeclaration won't work
	safe, rule := true, 60 //Redeclaring safe variable
	//safe := true this would not compile!

	fmt.Println(safe, rule)

	const (
		min int = 1
		max     //max gets stuff from above, it's the same as 'max int = 1'
	)
}

func simple() {
	fmt.Println("just a printing test")

	//Println accepts multiple inputs
	fmt.Println(1, 2, 3, 4, 5, 6)
	fmt.Println(true, false)
}

//Exported : Functions that start with uppercase are exported
//so if another file import this package, this function can be used
func Exported() {
	fmt.Println("Functions that start with uppercase are exported")
}

func converting() {
	speed := 100
	force := 2.5

	speed = speed * int(force) //2.5 becomes 2
	fmt.Println(speed)

	speed = int(float64(speed) * force)
	fmt.Println(speed)

	var apple int
	var orange int32

	apple = int(orange)
	orange = int32(apple)

	fmt.Println(apple, orange)
}

func slices() {
	var Args []string //A slice of strings
	Args = os.Args
	fmt.Println(len(Args), Args)
}

//An operation between int and float will produce float type
func arithmetic() {
	ratio := 1.0 / 10
	fmt.Printf("%.60f\n", ratio)

	//Trickier stuff
	var test float64 = 3 / 2 //this will produce 1
	fmt.Println(test)
	//What happens behind the scenes
	test = float64(int(3) / int(2))
	fmt.Println(test)

	//If we want the correct result
	test = 3.0 / 2
	fmt.Println(test)
}

func strings() {
	//srting literal
	var inline = "eita"

	//raw string literal
	var multiline = `
		DADSADASDAS
		DASDSADSADSADSADAS
		VXCVXCVXCVXCVX
	`

	fmt.Println(inline, multiline)

	fmt.Println("C:\\Intel\\Profiles\\IntelGraphicsProfiles\\test.exe")
	//This works because Go does not process what is inside raw string literal
	fmt.Println(`C:\Intel\Profiles\IntelGraphicsProfiles\test.exe`)

	//Using strconv
	i := 1
	fmt.Println("vish" + strconv.Itoa(i))
	fmt.Println(strconv.FormatBool(true) + " " + strconv.FormatBool(false))

	//len(...) functions does not count how many characters
	//there is in a string, but how many bytes. If there is
	//non-english chars, it will not work as expected. Best way is:
	fmt.Println(utf8.RuneCountInString("gsfgdjhfgsjd")) //rune can represent english and non-english
}

func definedTypes() {
	//We are creating a new Defined Type based on int32 predeclared type
	//int32 is the underlying type in that case
	//int32 is the real type! Duration is just a name
	type Duration int32

	var ms int32 = 1000
	var ns Duration

	//they are still different, we cannot do ms = ns, but we can:
	ns = Duration(ms)
	ms = int32(ns)

	fmt.Println(ms, ns)

	//Every Go type has an underlying type!
	//Example for int64: type int64 int64
	//As int64 is also a predeclared type, it can be defined by itself

	//ATTENTION: there is no type-hierarchy in Go, so
	type stuff Duration //still has underlying type = int32
}

//IOTA is a int generator
func iotaExample() {
	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)

	fmt.Println(EST, MST, PST)
}

func errorHandling() {
	i, err := strconv.Atoi("sa")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	//We can do better with short statements
	//This is called a 'Short If'
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("No error here! " + strconv.Itoa(n))
	}
}

//Comparing with Java, Go has the break behaviour by default
//If we need to avoid the break we can add the fallthrough
func switchExample() {
	city := "Wroclaw"

	switch city {
	case "Warsaw", "Krakow":
		fmt.Println("Warsaw or Krakow")
	default:
		fmt.Println("None")
	}

	i := 10

	//We can switch like this too
	switch {
	case i > 0:
		fmt.Println("bigger")
		fallthrough
	case i < 0:
		fmt.Println("smaller")
	default:
		fmt.Println("zero")
	}

	//Short Switch
	switch j := 10; true {
	case j > 0:
		fmt.Println("bigger")
		fallthrough
	case j < 0:
		fmt.Println("smaller")
	default:
		fmt.Println("zero")
	}
}

//for is the only loop statement in Go
func loops() {
	//Simple example
	for i := 0; i <= 5; i++ {
		fmt.Println("for " + strconv.Itoa(i))
	}

	j := 0

	//Another one
	for ; j <= 5; j++ {
		fmt.Println("another for " + strconv.Itoa(j))
	}

	//Infinite loop
	sum := 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}

	//Looping over a slice
}
