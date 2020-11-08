# GoLang Tutorial
This Is Where Anybody can Learn Go Basics!
Lets Do it One By One !

# Run Your Program
use`go run hello_world.go`command to run a golang program.

# Packages And  Declarations
We have to declare packages that has to be executed alongside. Every Packages has a Path and A Name Assiciated with.Packages are declared on top of the code.

# Operartors in Go
There are many operators in go like in anyother language.

    > Arithmetic Operators
    > Assignment Operators
    > Comparison Operators
    > Logical Operators
    > Bitwise Operators
* Other Operators *
Concatenation using + or += operators, & for address of, * for pointer to, <- for Receive operator.
# Escape Sequences
Special string comrise a backlash (\) followed by a letter or combination of digits for controlling output on monitor. For instance :  newline(\n), carriage returns(\r),tabs,(\t) and non printing characters in Screen.
    
    > \\        for \ character
    > \'        for ' character
    > \'        for ' character
    > \?        for ? character
    > \a        alert or bell
    > \b        for backspace
    > \f        for form feed
    > \n        for new line
    > \r        for carriage return
    > \t        for horizontal tab
    > \v        for vertical tab
    > \o        octal number of one to three digits
    > \xhh...   hexadecimal number of one or more digits

# Decisions
    > if statements
    > if else statements
    > Nested if statements
    > Switch statements
    > Select statements
    
Follow a program in code section for more in details.

# Loops
Loop is a programming structure that repeats a sequence of instructions until a specific condition is met. Go employs for loops and nested for loops for this purpose.
	
	for initialization; condition; increment {
		// loop body
	}

## Loop Control
Controlling loops are empoyed using 
   
    > break 
    > continue
    > goto statements
    
# Type Casting
Type conversion is a way to convert a variable from one data type to another data type. For example, if you want to store a long value into a simple integer then you can type cast long to int. You can convert values from one type to another using the cast operator. Its syntax is as follows −

`type_name(expression)`


# Functions

syntax:

	func function_name( [parameter list] ) [return_types]
	{
	   body of the function
	}
A function definition in Go programming language consists of a function header and a function body. Here are all the parts of a function −

**Func** − It starts the declaration of a function.

**Function Name** − It is the actual name of the function. The function name and the parameter list together constitute the function signature.

**Parameters** − A parameter is like a placeholder. When a function is invoked, you pass a value to the parameter. This value is referred to as actual parameter or argument. The parameter list refers to the type, order, and number of the parameters of a function. Parameters are optional; that is, a function may contain no parameters.

**Return Type** − A function may return a list of values. The return_types is the list of data types of the values the function returns. Some functions perform the desired operations without returning a value. In this case, the return_type is the not required.

**Function Body** − It contains a collection of statements that define what the function does.

# Variadic Functions
The function that called with the varying number of arguments is known as variadic function. Or in other words, a user is allowed to pass zero or more arguments in the variadic function. _fmt.Printf_ is the example of the variadic function, it required one fixed argument at the starting after that it can accept any number of arguments. e.g `variadic.go`

# Recursion
Recursion is the process of repeating items in a self-similar way. The same concept applies in programming languages as well. If a program allows to call a function inside the same function, then it is called a recursive function call.
`recursion.go`

# Closures
Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it. e.g `closure.go`

# Pointers
Every variable is a memory location and every memory location has its address defined which can be accessed using ampersand (&) operator, which denotes an address in memory.A pointer is a variable whose value is the address of another variable, i.e., direct address of the memory location. Like any variable or constant, you must declare a pointer before you can use it to store any variable address.

	syntax :
	`var var_name *var-type`

	var ip *int        /* pointer to an integer */
	var fp *float32    /* pointer to a float */
	
e.g : `pointers.go`

# Nil Pointer
Go compiler assign a Nil value to a pointer variable in case you do not have exact address to be assigned. This is done at the time of variable declaration. A pointer that is assigned nil is called a nil pointer.

The nil pointer is a constant with a value of zero defined in several standard libraries.
e.g. : `nil_pointer.go`

On most of the operating systems, programs are not permitted to access memory at address 0 because that memory is reserved by the operating system. However, the memory address 0 has special significance; it signals that the pointer is not intended to point to an accessible memory location. But by convention, if a pointer contains the nil (zero) value, it is assumed to point to nothing.

To check for a nil pointer 

	if(ptr != nil)     /* succeeds if p is not nil */
	if(ptr == nil)    /* succeeds if p is null */

# Array Of Pointers
?????????????????????????????????????????????????

# Structures 

Go arrays allow you to define variables that can hold several data items of the same kind. Structure is another user-defined data type available in Go programming, which allows you to combine data items of different kinds.

Structures are used to represent a record

syntax : 

definition :

	type struct_variable_type struct {
	   member definition;
	   member definition;
	   ...
	   member definition;
	}

declaration of stucture variables :

	variable_name := structure_variable_type {value1, value2...valuen}
	
	
e.g. : `structures.go`

# Structures as Function Arguments

You can pass a structure as a function argument in very similar way as you pass any other variable or pointer. e.g. : `structures_as_function_arguments.go`
