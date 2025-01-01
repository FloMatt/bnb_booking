package forms

type errors map[string][]string

// Add appends a new error message to the specified field in the errors map.
//
// Parameters:
//   - field: The name of the field to which the error message is associated.
//   - message: The error message to be added for the specified field.
//
// The function does not return any value. It modifies the errors map in-place.
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get retrieves the first error message associated with the specified field.
//
// Parameters:
//   - field: The name of the field for which to retrieve the error message.
//
// Returns:
//
//	A string containing the first error message for the specified field.
//	If no errors exist for the field, an empty string is returned.
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
