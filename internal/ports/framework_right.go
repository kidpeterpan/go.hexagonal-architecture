package ports

type DbPort interface {
	AddToHistory(answer int32, operation string) error
	CloseDbConnection() error
}
