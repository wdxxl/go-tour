package main

type myError struct {
	Msg string
}

func (m *myError) Error() string {
	return m.Msg
}

func raiseErrorWhenBadThingsHappens(bad bool) error {
	if bad {
		return &myError{Msg: "this is bad!"}
	}
	return nil
}
