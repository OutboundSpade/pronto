>This page is a work in progress. Things may be added/removed/changed
- [Variables](#variables)
	- [Types](#types)
		- [Modifiers](#modifiers)
	- [Declaration](#declaration)
- [Operators](#operators)
	- [Math](#math)
	- [Boolean](#boolean)
- [Conditionals](#conditionals)

---

## Variables

### Types
Pronto is statically typed with the following built-in types:

- `num` - Number (Integer/Decimal/Float)
- `str` - String (denoted by `""` or `''`)
- `obj` - Object (contains one or more of any type, including another object)
- `bool` - Boolean (`true` or `false`)

#### Modifiers
Currently, the only modifier is `list` which indicates list of a certain type
All modifiers are placed before the type

See [Declaration](#declaration) for more information

### Declaration

A variable can be declared by the following:
```
<modifier?> <type> <name> = <value>
```
- `<modifier?>` (optional) is a valid modifier (see [Modifiers](#modifiers))
- `<type>` is a valid type (see [Types](#types))
- `<name>` is a valid identifier (matches the pattern `[_a-zA-Z][_a-zA-Z0-9]{0,255}`)
- `<value>` is a valid value for the specified type (see [Types](#types))

>Note: This may be changed in the future as type inferring is still on the table

Examples:
<details>
	<summary>Declaring a number</summary>

	num a = 5.1
	
</details>

<details>
	<summary>Declaring a string</summary>

	str a = "Hello"
	
</details>
<details>
	<summary>Declaring a boolean</summary>

	bool a = true
	
</details>
<details>
	<summary>Declaring an object</summary>

	obj a = {}
	
</details>
<details>
	<summary>Declaring an object with properties</summary>

	obj person1 = {
		str name: "Dwight",
		num age: 51,
		obj title: {
			str name: "Asst. to the Regional Manager",
			num rank: 2
		}
	}
	
</details>

## Operators

### Math

- `+` add
- `-` subtract
- `*` multiply
- `/` divide
- `^` power (`3 ^ 2` is `3²`)
- `^^` root (`9 ^^ 2` is `√9`)

### Boolean

- `and` - (`&&` in other languages)
- `or` - (`||` in other languages)
- `is` or `==` - equality (`==` in other languages)
- `>` - greater than
- `>=` - greater than or equal to (`≥`)
- `<` - less than
- `<=` - less than or equal to (`≤`)


## Conditionals
