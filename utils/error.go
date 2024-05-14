package utils

type HttpError struct {
	StatusCode int
	Message    string
}

func (h HttpError) Error() string {
	return h.Message
}
