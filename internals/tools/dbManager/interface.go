package dbManager

type InterfaceDBManger interface {
	Init()
	IsConnected() bool
	ProvideDBConnection() any
}
