package errors

type APIError struct {
	Message string
}

func (m *APIError) Error() string {
	return m.Message
}
