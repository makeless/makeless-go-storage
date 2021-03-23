package basic

import (
	"github.com/spf13/afero"
	"os"
	"sync"
)

type Storage struct {
	fs afero.Fs
	*sync.RWMutex
}

func (storage *Storage) getFs() afero.Fs {
	storage.Lock()
	defer storage.Unlock()

	return storage.fs
}

func (storage *Storage) setFs(fs afero.Fs) {
	storage.RLock()
	defer storage.RUnlock()

	storage.fs = fs
}

func (storage *Storage) Init() error {
	storage.setFs(afero.NewOsFs())
	return nil
}

func (storage *Storage) Write(filepath string, data []byte, perm os.FileMode) error {
	return afero.WriteFile(storage.getFs(), filepath, data, perm)
}

func (storage *Storage) Read(filepath string) ([]byte, error) {
	return afero.ReadFile(storage.getFs(), filepath)
}

func (storage *Storage) Remove(filepath string) error {
	return storage.getFs().Remove(filepath)
}

func (storage *Storage) Exists(filepath string) (bool, error) {
	return afero.Exists(storage.getFs(), filepath)
}
