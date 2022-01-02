>This page is a work in progress. Things may be added/removed/changed
- [Basic variable usage](#basic-variable-usage)
	- [Numbers](#numbers)
	- [Strings](#strings)
	- [Lists](#lists)
- [Conditional statements](#conditional-statements)
	- [Basic if](#basic-if)
	- [if else](#if-else)

---

### Basic variable usage
#### Numbers
```
num a = 1 // sets a to 1
num b = a // sets b to a (1)
b++ //adds 1 to b

num c = b + a // sets c to a (1) + b (2)
log c //logs c (3)
```
#### Strings
```
str name = "Hello"
int nameLength = len name

```
#### Lists
>Lists and Arrays are the same thing in pronto
```
list num a = [1,2,3]
list str b = ["Michael","Dwight","Jim","Ryan"]
list obj c = [{str text:"hi"},{str text:"hello"}]
log a // -> 1, 2, 3
log len(a) // -> 1, 2, 3
```

### Conditional statements

#### Basic if
```
if a is 1 {
	//do stuff
}
```
#### if else
```

```

