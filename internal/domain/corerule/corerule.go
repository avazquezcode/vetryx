package corerule

import "fmt"

// IsTrue is the rule used to determine whether something should be evaluated to true or false.
func IsTrue(value interface{}) bool {
	if value == nil {
		return false
	}

	boolean, isBoolean := value.(bool)
	if isBoolean {
		return boolean
	}

	return true
}

// PrintableValue converts an interface into a printable value.
func PrintableValue(value interface{}) string {
	if value == nil {
		// in this language, nil is represented as "null"
		return "null"
	}
	return fmt.Sprintf("%v", value)
}

// IsEqual is the rule used to determine whether two values are equal.
func IsEqual(a interface{}, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil {
		return false
	}

	return a == b
}
