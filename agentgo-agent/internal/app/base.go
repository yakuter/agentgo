package app

type Executer interface {
	Execute() (string, error)
}
