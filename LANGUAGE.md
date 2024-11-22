# Language

## Intro

This file intends to document in a summarised way the syntax of Vetryx, and some topics related to its grammar.

## Characteristics of the language
- Interpreted
- Dynamically typed
- High level
- Imperative

## Types

This language only support basic types:

| Operator | Description |
| ----------- | ----------- |
| string | "hello world" |
| number | Eg: 1. *Note*: _(*All numbers are floats for now*)_ |
| bool | true / false |
| null | null value |

## Operators

### Arithmetics

| Operator | Description |
| ----------- | ----------- |
| + | Sum two numbers, _or concat strings_ |
| - | Subtract two numbers |
| * | Multiply two numbers |
| / | Divide two numbers |
| % | Modulus between two numbers |

### Comparators

| Operator | Description |
| ----------- | ----------- |
| == | Equal |
| <> | Different |
| > | Greater |
| >= | Greater or Equal |
| < | Lower |
| <= | Lower or Equal |

### Unary

| Operator | Description |
| ----------- | ----------- |
| ! | Eg: !false => true |
| - | Negates a number. Eg: -(-1) => 1 |


### Logical Operators

| Operator | Description |
| ----------- | ----------- |
| && | AND |
| &#124;&#124; | OR |

## In-built Functions

| Operator | Description |
| ----------- | ----------- |
| print | prints anything to the stdout |
| clock | returns the current timestamp in nanoseconds (based on the clock) |
| sleep(X) | add a delay of "X" ms to the execution of the program |
| min(X, Y) | returns min |
| max(X, Y) | returns max |

## Reserved Words

| Word | 
| ----------- |
| while |
| break |
| continue |
| if |
| else |
| dec |
| fn |
| return |
| print |
| null |
| true |
| false |

## Comments

You can add comments by using `#`. Example:

```python
# this line will be ignored
```

## Grouping

You can group expressions using parentheses. Example:

```python
print (1+3) * 2; # will print 8, first it resolves the content of parentheses, then the multiplication.
```

## Variables

### Declaration

A variable can be declared empty (without assignment), and in such case, will be `null` by default:

```python
dec a;
print a; # prints null
```

A variable can be declared and assigned in the same line:

```python
dec a = 1;
print a; # prints 1
```

A variable can also be declared and assigned using the "short declarator" (similar as in _Go_):

```python
a := 1;
print a; # prints 1
```

### Assignment

You can assign a value to an existing variable:

```python
dec a;
a = 1;
print a; # prints 1
```

📌 *Important*: If the variable is not declared before assignment, the interpreter will throw an error.

## If

The syntax for the if condition is:

```python
if "a" == "a" {
    print 1; # will print 1
}

# You can also wrap the condition with parentheses:

if ("a" == "a") {
    print 2; # will print 2
}
```

If you want to add an else condition, you can also do it:

```python
if ("a" <> "a") {
    print 2; 
} else {
    print 1; # will print 1 in this case
}
```

## While

The syntax for the while loop is:

```python
dec a = 10;
while a < 20 {
    print a; # will print the value of "a" in the current iteration
    a = a + 1;
}
```

You can also wrap the condition with parentheses:

```python
dec a = 10;
while (a < 20) {
    print a; # will print the value of "a" in the current iteration
    a = a + 1;
}
```

You can use break to exit the loop:

```python
dec a = 10;
while (a < 20) {
    break; # will exit the loop inmediately
}
```

You can use continue to move to next iteration of the loop:

```python
dec a = 10;
while (a < 20) {
    a = a + 1;
    if a == 15 {
        continue; # in this case the 15 will not be printed to stdout
    }
    print a;
}
```

## Functions

### Declaration

```python
# Without parameters:
fn a() {
    return "x"; # note: if you ommit the return, it will return null by default
}
```

```python
# With parameters:
fn a(b, c) {
    return b + c;
}
```

### Call

```python
fn a(b, c) {
    return b + c;
}

print a(1, 2); # will print 3 
```

If you want to assign the value to a variable, first you need to declare it, then assign it. Example:

```python
dec x;
x = a(1, 2);
print x; # will print 3
```

### Closures

Closures are supported in the language.

```python
fn buildCounter() {
  i := 0;
  fn count() {
    i = i + 1;
    print i;
  }

  return count
}

dec counter;
counter = buildCounter();

counter(); # Prints 1
counter(); # Prints 2
counter(); # Prints 3
```
