package utils

type ResponseError struct {
	StatusCode string
	Body       string
	Err        string
}

func (r ResponseError) Error() string {
	return r.Err
}
