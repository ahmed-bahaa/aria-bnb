package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Valid returns true if there no erros, otherwise return false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Form create a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initiallize a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form fields in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x != "" {
		return true
	}
	f.Errors.Add(field, "this field cannot be blank")
	return false
}

// Required check for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field Cannot be Blank")
		}
	}
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}
