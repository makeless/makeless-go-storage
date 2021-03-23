package basic

import (
	"github.com/spf13/afero"
	"sync"
	"testing"
)

func TestStorageWriteRead(t *testing.T) {
	var storage = &Storage{
		fs:      afero.NewMemMapFs(),
		RWMutex: new(sync.RWMutex),
	}

	var tests = []struct {
		storage  *Storage
		filepath string
		data     []byte
	}{
		{
			storage:  storage,
			filepath: "test.dat",
			data:     []byte("test"),
		},
		{
			storage:  storage,
			filepath: "test/test-1/test-2/test.dat",
			data:     []byte("test"),
		},
	}

	for i, test := range tests {
		var err error
		var data []byte

		if err = test.storage.Write(test.filepath, test.data, 0644); err != nil {
			t.Errorf("%d: error: %s", i, err.Error())
			return
		}

		if data, err = test.storage.Read(test.filepath); err != nil {
			t.Errorf("%d: error: %s", i, err.Error())
			return
		}

		if string(test.data) != string(data) {
			t.Errorf("%d: error: %s", i, "bytes not equal")
		}
	}
}
