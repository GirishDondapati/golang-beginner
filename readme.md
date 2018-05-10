
Go Language Basics explanations and related example programs.
===============================================================

1. Packages
-----------
Every Go program is made up of packages. Programs start running in package main.
This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance,
the "math/rand" package comprises files that begin with the statement package rand.

Note: the environment in which these programs are executed is deterministic, so each time you run the example program rand.
In this will return the same number. (To see a different number, seed the number generator; see rand.Seed. Time is constant in the playground,
so you will need to use something else as the seed.)

Run in terminal: go run packages_efine


2. Imports
----------
This code groups the imports into a parenthesized, "factored" import statement.
You can also write multiple import statements, like:

import "fmt"
import "math"

But it is good style to use the factored import statement.

Run in terminal: go run import_types


3) Functions
----------

A function can take zero or more arguments.In given example, add takes two parameters of type int.
 (For more about why types look the way they do, see the article on Go's declaration syntax.)
Run in terminal: go run create_function
Declaring function in go

Run in terminal: go run create_function


4) Variables with initializers
------------------------------
A var declaration can include initializers, one per variable.
If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

Go's basic types are
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.
The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

Run in terminal: go run declare_variables


5. Variables and Constants declaration
--------------------------------------

Basic types variable types

	bool
	string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // alias for uint8
	rune // alias for int32
    		 // represents a Unicode code point
	float32 float64
	complex64 complex128

The example shows variables of several types, and also that variable declarations may
 be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits
 wide on 64-bit systems. When you need an integer value you should use int unless you 
have a specific reason to use a sized or unsigned integer type.

Constants declaration:

Constants are declared like variables, but with the const keyword.
Constants can be character, string, boolean, or numeric values.
Constants cannot be declared using the := syntax.

Run in terminal as: go run variables_constants




