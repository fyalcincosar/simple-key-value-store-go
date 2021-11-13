package usecases

type Logger interface {
	LogError(string, ...interface{})
	LogHttpServer(string, ...interface{})
}
