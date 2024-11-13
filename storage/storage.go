package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Storage[T any] struct {
	FileName string
	saveCh   chan T
	closeCh  chan struct{}
	wg       sync.WaitGroup
}

func NewStorage[T any](fileName string) *Storage[T] {
	storage := &Storage[T]{
		FileName: fileName,
		saveCh:   make(chan T),
		closeCh:  make(chan struct{}),
	}

	storage.wg.Add(1)
	go storage.saveWorker()

	return storage
}

func (s *Storage[T]) Save(data T) {
	s.saveCh <- data // send data to channel
}

func (s *Storage[T]) Close() {
	close(s.saveCh)
	s.wg.Wait() // wait for saveWorker to finish
	close(s.closeCh)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", s.FileName, err)
	}

	if err := json.Unmarshal(fileData, data); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return nil
}

func (s *Storage[T]) saveWorker() {
	defer s.wg.Done()
	for data := range s.saveCh {
		fileData, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			fmt.Println("failed to marshal data:", err)
			continue
		}
		err = os.WriteFile(s.FileName, fileData, 0644)
		if err != nil {
			fmt.Println("failed to write file:", err)
		}
	}
}
