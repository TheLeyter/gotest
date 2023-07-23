package errors

type DBError struct {
	Message string
}

func (m *DBError) Error() string {
	return m.Message
}
