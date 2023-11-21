package ustring

import "fmt"

var resultStr string

// RewriteSymbol returns the updated string with a modified element.
// The symbol changes the element with the index specified in the function call
func RewriteSymbol(str string, index int, value string) (string, error) {
	if index < 0 || index > len(str)-1 {
		return "", fmt.Errorf("error: the index [%v] does not exist or is not in the given string", index)
	}

	resultStr = str[:index] + value + str[index+1:]
	return resultStr, nil
}

// AddSymbol returns an updated string with the new element.
// The new element is placed on the index specified in the function call.
func AddSymbol(str string, index int, value string) (string, error) {
	if index < 0 || index > len(str) {
		return "", fmt.Errorf("error: the index [%v] cannot be found or it exceeds the length of the string", index)
	}

	resultStr = str[:index] + value + str[index:]
	return resultStr, nil
}

// RemoveSymbol returns an updated string with the deleted item.
func RemoveSymbol(str string, index int) (string, error) {
	if index < 0 || index > len(str) {
		return "", fmt.Errorf("error: the index [%v] cannot be found or it exceeds the length of the string", index)
	}

	resultStr = str[:index] + str[index+1:]
	return resultStr, nil
}
