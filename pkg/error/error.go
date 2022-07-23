package errors

type CustomError struct {
	Code        int
	Message     string
	InfoContext interface{} `json:"info_context,omitempty"`
}

func (err CustomError) Error() string {
	return err.Message
}

func ToCustomError(err error) CustomError {
	if err, ok := err.(CustomError); ok {
		return err
	}
	return CustomError{
		Code:    500,
		Message: err.Error(),
	}
}
