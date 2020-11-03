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
