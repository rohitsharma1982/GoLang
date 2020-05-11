package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	/*
		If empty name, returns the string "One for you, one for me."
		else returns "One for name, one for me."
	*/
	/*
		switch name {
			case "":
				return "One for you, one for me."
			default:
				return "One for " + name + ", one for me."
		}
	*/
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
