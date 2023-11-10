package store

import (
	"encoding/csv"
	"os"
)

type Store struct {
	path string
}

func New(path string) *Store {
	return &Store{
		path: path,
	}
}

func (s *Store) Read() ([][]string, error) {
	f, err := os.Open(s.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	return reader.ReadAll()
}

func (s *Store) Write(data []string) error {
	file, err := os.OpenFile(s.path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	if err := w.Write(data); err != nil {
		return err
	}
	w.Flush()

	return nil
}
