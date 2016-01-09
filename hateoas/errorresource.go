package hateoas

// ErrorResource satisfies both the building error interface
// and the Resource interface
type ErrorResource struct {
	error error
}

// NewErrorResource creates an ErrorResource from the
// provided error
func NewErrorResource(err error) Resource {
	return &ErrorResource{
		error: err,
	}
}

// Links satisfies the Resource interface.
func (e *ErrorResource) Links() *Links {
	return nil
}

func (e *ErrorResource) Error() string {
	return e.error.Error()
}

// MarshalJSON satisfies the MarshalJSON interface
func (e *ErrorResource) MarshalJSON() ([]byte, error) {
	return []byte(`{"error":"` + e.Error() + `"}`), nil
}
