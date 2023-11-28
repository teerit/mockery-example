package example_project

import "fmt"

// Store provides the interface for storing keyed data
type Store interface {
	Write(key string, data []byte) error
	Read(key string) ([]byte, error)
}

// FileStore is an implementation of the Store interface which stores data in files
type FileStore struct {
	directoryPath string
}

func (f *FileStore) Write(key string, data []byte) error {
	// Implementation
	fmt.Println(f.directoryPath)
	return nil
}

func (f *FileStore) Read(key string) ([]byte, error) {
	// Implementation
	fmt.Println(f.directoryPath)
	return nil, nil
}

func NewFileStore(path string) (Store, error) {
	// pointer to concrete type
	return &FileStore{directoryPath: path}, nil
}
