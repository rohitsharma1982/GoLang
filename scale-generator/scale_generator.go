package scale

import "strings"

var stepSize = map[string]int{ // Map relating interval to number of notes we need to jump to create the scale
	"m": 1,
	"M": 2,
	"A": 3,
}

func Scale(tonic string, interval string) []string {
	sharpScale := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flatScale := []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

	result := make([]string, 0)

	if interval == "" { // An empty interval cover all the notes in the particular scale, hence replacing it with consecutive steps
		interval = "mmmmmmmmmmmm"
	}

	if isInSharps(tonic) { // Checking if tonic is in Sharps or Flats. For simplicity, C major and a minor are added to Sharps.
		tonic = convertToUpper(tonic) // Sharp or Flat scale does not have minor notes, hence converting d# to D# and bb to Bb
		result = append(result, tonic)
		noteIndex := getNoteIndex(tonic, sharpScale) // Getting index of starting tonic in the respective scale

		for _, character := range interval {
			noteIndex += stepSize[string(character)]
			noteIndex %= len(sharpScale) // In case the index goes out of bounds, using % operator to start the counting from 0
			if sharpScale[noteIndex] != tonic {
				result = append(result, sharpScale[noteIndex])
			}
		}
	} else {
		tonic = convertToUpper(tonic)
		result = append(result, tonic)
		noteIndex := getNoteIndex(tonic, flatScale)

		for _, character := range interval {
			noteIndex += stepSize[string(character)]
			noteIndex %= len(flatScale)
			if flatScale[noteIndex] != tonic {
				result = append(result, flatScale[noteIndex])
			}
		}
	}

	return result
}

func convertToUpper(tonic string) string {
	if len(tonic) == 1 {
		return strings.ToUpper(tonic)
	} else {
		return strings.ToUpper(string(tonic[0])) + string(tonic[1])
	}
}

func isInSharps(tonic string) bool {
	match := false
	sharpNotes := []string{"C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#"}

	for itr := range sharpNotes {
		if tonic == sharpNotes[itr] {
			match = true
			break
		}
	}

	return match
}

func getNoteIndex(tonic string, scaleArray []string) int {
	found := 0

	for itr := range scaleArray {
		if tonic == scaleArray[itr] {
			found = itr
			break
		}
	}

	return found
}
