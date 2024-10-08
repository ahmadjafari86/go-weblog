package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"required": "The field is required",
		"email":    "The field must be filled with valid email address",
		"min":      "Should be more than the limit",
		"max":      "Should be less than the limit",
	}
}
