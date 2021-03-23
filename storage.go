package makeless_go_storage

type Storage interface {
	Init() error
	Write(filepath string, data []byte) error
	Read(filepath string) ([]byte, error)
	Remove(filepath string) error
	Exists(filepath string) (bool, error)
}
