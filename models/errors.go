package models

type MyError struct {
	Err string `json:"error"`
}

func (me MyError) Error() string {
	return me.Err
}

var (
	ErrorInvalidEmailOrPassword = MyError{
		Err: "invalid email / password",
	}

	ErrorInvalidToken = MyError{
		Err: "invalid token",
	}

	ErrorNotAuthorized = MyError{
		Err: "Not Authorized",
	}

	ErrorNotFound = MyError{
		Err: "Not Found",
	}

	ErrorForbiddenAccess = MyError{
		Err: "Forbidden Access",
	}
)
