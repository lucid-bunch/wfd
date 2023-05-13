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
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (s *Store) Write(data []string) error {
	file, err := os.OpenFile(s.path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()
	if err := w.Write(data); err != nil {
		return err
	}

	return nil
}
