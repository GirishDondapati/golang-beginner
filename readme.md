
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

