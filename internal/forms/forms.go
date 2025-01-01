package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct that embeds a url.Values object and an errors map.
type Form struct {
	url.Values
	Errors errors
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New creates a new Form instance with the provided data and initializes an empty errors map.
//
// The function takes a url.Values object as a parameter, which represents the form data.
// It returns a pointer to a Form struct, which contains the provided data and an empty errors map.
//
// Example usage:
//
//	data := url.Values{
//		"username": {"john_doe"},
//		"email":    {"john@example.com"},
//	}
//	form := New(data)
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks if the specified form fields are empty. If any of the fields are empty, an error message is added to the errors map.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		x := f.Get(field)
		if strings.TrimSpace(x) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength checks if the specified form field has a minimum length. If the field does not meet the minimum length requirement, an error message is added to the errors map.
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	X := r.Form.Get(field)
	if len(X) < length {
		f.Errors.Add(field, fmt.Sprintf("Field must be at least length of %d characters", length))
		return false
	}
	return true
}

// Has checks if a form field is empty. If the field is empty, an error message is added to the errors map.

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {

		return false
	}
	return true
}

// IsEmail checks if the specified form field is a valid email address. If the field is not a valid email address, an error message is added to the errors map.
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
