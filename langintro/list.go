package langintro

// MarioCharacters returns the list of Mario Characters
func MarioCharacters() []string {
	// Slices are actually arrays without a fixed length
	// Here we are returning a slice and not an array, slices actually use arrays under the hood
	return []string{"Mario", "Luigi", "Bowser", "Peach"}
}

// Notes:
// Arrays in Go are exactly similar to Arrays in C
// Appending elements to a slice actually returns a new slice, it doesn't modify the original slice
// Slice Ranges is actually getting a range of elements from an existing slice and storing them in a new slice
