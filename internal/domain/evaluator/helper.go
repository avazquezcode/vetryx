package evaluator

func isNumber(operand ...interface{}) bool {
	for _, item := range operand {
		if _, ok := item.(float64); !ok {
			return false
		}
	}

	return true
}

func isString(operand ...interface{}) bool {
	for _, item := range operand {
		if _, ok := item.(string); !ok {
			return false
		}
	}

	return true
}
