package errors

// String creates an error from a string
type String string

// Error implements error interface
func (s String) Error() string { return string(s) }

// String implements stringer interface
func (s String) String() string { return string(s) }

// Wrapper wraps an error and a cause
type Wrapper struct {
	Err     error
	Because string
}

// Wrap the cause of an error with the error
func Wrap(cause string, err error) error {
	if err == nil {
		return nil
	}
	return &Wrapper{
		Err:     err,
		Because: cause,
	}
}

// Error implements error interface
func (w *Wrapper) Error() string {
	return w.Because + w.Err.Error()
}

// Cause returns the cause of the error
func (w *Wrapper) Cause() string {
	return w.Because
}
