// package main

// import "fmt"

// func main() {
// 	// 1. Declaring Variables
// 	//    Variables can be declared using the 'var' keyword, followed by the variable name and type.
// 	var age int       // Declaration of an integer variable named 'age'
// 	var name string   // Declaration of a string variable named 'name'
// 	var isStudent bool // Declaration of a boolean variable named 'isStudent'

// 	// 2. Initializing Variables
// 	//    Variables can be initialized during declaration or later.
// 	age = 30           // Assigning a value to 'age'
// 	name = "Alice"     // Assigning a value to 'name'
// 	isStudent = true   // Assigning a value to 'isStudent'

// 	fmt.Println("--- Declaring and Initializing Variables ---")
// 	fmt.Printf("Name: %s, Age: %d, Is Student: %t\n", name, age, isStudent)

// 	// 3. Short Variable Declaration (Type Inference)
// 	//    The ':=' operator can be used inside functions to declare and initialize
// 	//    a variable, and Go will infer the type based on the assigned value.
// 	city := "New York" // 'city' is inferred as a string
// 	zipCode := 10001   // 'zipCode' is inferred as an int
// 	gpa := 3.8         // 'gpa' is inferred as a float64

// 	fmt.Println("\n--- Short Variable Declaration ---")
// 	fmt.Printf("City: %s, Zip Code: %d, GPA: %.1f\n", city, zipCode, gpa)

// 	// 4. Multiple Variable Declaration
// 	//    You can declare multiple variables of the same type in a single line.
// 	var x, y int = 10, 20
// 	fmt.Printf("\nx: %d, y: %d\n", x, y)

// 	//    Or declare multiple variables of different types in a block.
// 	var (
// 		firstName string = "Bob"
// 		lastName  string = "Smith"
// 		height    float64 = 1.75
// 	)
// 	fmt.Printf("Name: %s %s, Height: %.2f meters\n", firstName, lastName, height)

// 	// 5. Basic Data Types

// 	//    Integers: int, int8, int16, int32, int64 (signed)
// 	//              uint, uint8, uint16, uint32, uint64, uintptr (unsigned)
// 	var numInt int = 42
// 	var numInt8 int8 = 127
// 	var numUint uint = 100
// 	fmt.Println("\n--- Integer Types ---")
// 	fmt.Printf("int: %d, int8: %d, uint: %d\n", numInt, numInt8, numUint)

// 	//    Floating-Point Numbers: float32, float64
// 	var pi float32 = 3.14
// 	var e float64 = 2.71828
// 	fmt.Println("\n--- Floating-Point Types ---")
// 	fmt.Printf("float32 (pi): %.2f, float64 (e): %.5f\n", pi, e)

// 	//    Booleans: bool (true or false)
// 	var isActive bool = true
// 	var hasPermission bool = false
// 	fmt.Println("\n--- Boolean Types ---")
// 	fmt.Printf("Is Active: %t, Has Permission: %t\n", isActive, hasPermission)

// 	//    Strings: string (sequence of characters)
// 	var message string = "Hello, Go!"
// 	var greeting string = `This is a
// multi-line string
// using backticks.`
// 	fmt.Println("\n--- String Types ---")
// 	fmt.Println("Message:", message)
// 	fmt.Println("Greeting:", greeting)

// 	// 6. Zero Values
// 	//    If variables are declared but not explicitly initialized, they are
// 	//    automatically assigned a "zero value".
// 	var defaultInt int      // 0
// 	var defaultString string // ""
// 	var defaultBool bool    // false
// 	var defaultFloat float64 // 0.0
// 	fmt.Println("\n--- Zero Values ---")
// 	fmt.Printf("Default int: %d, Default string: %s, Default bool: %t, Default float: %.1f\n",
// 		defaultInt, defaultString, defaultBool, defaultFloat)

// 	// 7. Type Conversion
// 	//    Go is a strongly-typed language, so explicit type conversion is often needed.
// 	var intVal int = 10
// 	var floatVal float64 = float64(intVal) // Convert int to float64
// 	var anotherInt int = int(floatVal)     // Convert float64 back to int (truncates decimal)

// 	fmt.Println("\n--- Type Conversion ---")
// 	fmt.Printf("intVal: %d, floatVal: %.2f, anotherInt: %d\n", intVal, floatVal, anotherInt)

// 	var strNum string = "123"
// 	// To convert string to int, you'd typically use the "strconv" package:
// 	// import "strconv"
// 	// num, err := strconv.Atoi(strNum)
// 	// if err != nil { /* handle error */ }
// 	// fmt.Println("Converted string to int:", num)
// 	fmt.Printf("To convert string \"%s\" to int, use 'strconv.Atoi'\n", strNum)

// 	// 8. Constants
// 	//    Constants are declared with the 'const' keyword and cannot be changed.
// 	const PI float64 = 3.14159
// 	const GREETING string = "Hello, Constant!"
// 	const TRUTH = true // Type inference also works for constants

// 	fmt.Println("\n--- Constants ---")
// 	fmt.Printf("PI: %.5f, Greeting: %s, Truth: %t\n", PI, GREETING, TRUTH)

// 	// 9. Iota (for enumerations)
// 	//    'iota' is a predeclared identifier that represents successive untyped
// 	//    integer constants. It resets to 0 whenever the word 'const' appears.
// 	const (
// 		MONDAY = iota // 0
// 		TUESDAY       // 1
// 		WEDNESDAY     // 2
// 	)
// 	fmt.Println("\n--- Iota (Enumerations) ---")
// 	fmt.Printf("Monday: %d, Tuesday: %d, Wednesday: %d\n", MONDAY, TUESDAY, WEDNESDAY)

// 	const (
// 		_ = iota // 0 (blank identifier to discard the first value)
// 		KB = 1 << (10 * iota) // 1 << (10 * 1) = 1024
// 		MB = 1 << (10 * iota) // 1 << (10 * 2) = 1048576
// 		GB = 1 << (10 * iota) // 1 << (10 * 3) = 1073741824
// 	)
// 	fmt.Printf("KB: %d, MB: %d, GB: %d\n", KB, MB, GB)
// }