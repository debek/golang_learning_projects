## 1.

### Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” “y” and “z”

- a) 42
- b) James Bond"
- c) true

Now print the values stored in those variables using

- a single print statement
- multiple print statements

## 2.

### 1. Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).

- identifier “x” type int
- identifier “y” type string
- identifier “z” type bool

### 2. in func main

- print out the values for each identifier
- The compiler assigned values to the variables. What are these values called?

## 3.

### Using the code from the previous exercise,

### 1. At the package level scope, assign the following values to the three variables

- for x assign 42
- for y assign “James Bond”
- for z assign true

### 2. in func main

- use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
- print out the value stored by variable “s”

## 4.

### 1. Create your own type. Have the underlying type be an int.

### 2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword

### 3. in func main

- print out the value of the variable “x”
- print out the type of the variable “x”
- assign 42 to the VARIABLE “x” using the “=” OPERATOR
- print out the value of the variable “x”

## 5.

### Building on the code from the previous example

### 1. at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”

example:

```
type hotdog int
var x hotdog
var y int
```

### 2. in func main

a. this should already be done

- print out the value of the variable “x”
- print out the type of the variable “x”
- assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
- print out the value of the variable “x”

b. now do this

now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE

- then use the “=” operator to ASSIGN that value to “y”
- print out the value stored in “y”
- print out the type of “y”