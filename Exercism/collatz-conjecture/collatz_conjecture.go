package collatzconjecture

import "fmt"

func CollatzConjecture(input int) (int, error) {

	var steps int = 0
	var err error = nil

	if input == 0 {
		err = fmt.Errorf("Error: Zero is an error")
	} else if input < 0 {
		err = fmt.Errorf("Error: Negative value is an error")
	} else if input == 1 {
		steps = 0
	} else {
		for input > 1 {
			steps += 1

			if input%2 == 0 {
				input /= 2
			} else {
				input = 3*input + 1
			}
		}
	}

	return steps, err
}
