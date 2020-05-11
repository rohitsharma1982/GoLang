package bob

/*
The function Hey returns Bob's answer to anything spoken to him.
*/

/*
Bob is a lackadaisical teenager. In conversation, his responses are very limited.

Bob answers 'Sure.' if you ask him a question, such as "How are you?".

He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).

He answers 'Calm down, I know what I'm doing!' if you yell a question at him.

He says 'Fine. Be that way!' if you address him without actually saying
anything.

He answers 'Whatever.' to anything else.

Bob's conversational partner is a purist when it comes to written communication and always follows normal rules regarding sentence punctuation in English.
*/

import (
	"strings"
	"unicode"
)

// Hey returns Bob's response to the sentence addressed to him.
// The details of Bob's responses to given sentences is mentioned earlier.
func Hey(remark string) string {

	var reply string

	remark = strings.TrimSpace(remark) //Trimming the remark string of any empty spaces at the end or at the start

	if len(strings.Fields(remark)) == 0 { //Reply to empty comment

		reply = "Fine. Be that way!"

	} else if strings.HasSuffix(remark, "?") { //Reply to a question

		remark = strings.TrimSuffix(remark, "?")
		nocaps := true

		for _, character := range remark {
			if unicode.IsLetter(character) {
				if unicode.IsLower(character) {
					nocaps = true
					break
				} else {
					nocaps = false
				}
			} else {
				nocaps = true
			}
		}

		if nocaps {
			reply = "Sure."
		} else {
			reply = "Calm down, I know what I'm doing!"
		}

	} else { //Reply to a statement

		allcaps := false

		for _, character := range remark {
			if unicode.IsLetter(character) {
				if !unicode.IsUpper(character) {
					allcaps = false
					break
				} else {
					allcaps = true
				}
			}
		}

		if allcaps {
			reply = "Whoa, chill out!"
		} else {
			reply = "Whatever."
		}
	}

	return reply
}
