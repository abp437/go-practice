package langintro

func updateNameViaPointer(n *string) {
	*n = *n + " Pawar!"
}

// ModifyName accepts a name and then modifies it's value through a pointer
func ModifyName(n string) string {
	var tempPointer *string
	tempPointer = &n
	// This works as well
	// tempPointer := &n

	updateNameViaPointer(tempPointer)
	// This works as well
	// updateNameViaPointer(&n)
	return n
}

// Notes:
// Pointers are a data type in itself
// Pointers are generally used for data types which are pass by value, such as strings, arrays and structs.
// But the extensive usage is for a Struct.
// Poniters as usual are denoted by "*" in function signature, declaration and used with a "*" for dereferencing.
// Only while declaring a pointer or dereferencing it we need a "*", the normal usages don't require a "*" in the code.
// Pointers generally have the purpose of cross functional variables since the scope of variables is functional in most languages.
// Also, it has the advantage of direct manipulation of data without excessive passing of data through functions.
