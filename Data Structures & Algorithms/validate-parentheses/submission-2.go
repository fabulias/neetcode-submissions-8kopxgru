func isValid(s string) bool {
    myStack := []rune{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			myStack = append(myStack, c)
		} else {
			if len(myStack) == 0 {
				return false
			}
			lastInsertion := myStack[len(myStack)-1]
			isValid := false
			if lastInsertion == '(' && c == ')' {
				isValid = true
			}
			if lastInsertion == '{' && c == '}' {
				isValid = true
			}
			if lastInsertion == '[' && c == ']' {
				isValid = true
			}
			if !isValid {
				return isValid
			} else {
				myStack = myStack[:len(myStack) - 1]
			}
		}
	}
	if len(myStack) > 0 {
		return false
	}
	return true
}
