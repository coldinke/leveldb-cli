package db

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Manager struct {
	db   *leveldb.DB
	path string
}

type KeyValue struct {
	Key   string
	Value string
}

func NewManager(path string) (*Manager, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return &Manager{
		db:   db,
		path: path,
	}, nil
}

func (m *Manager) Close() error {
	return m.db.Close()
}

func (m *Manager) Get(key string) (string, error) {
	value, err := m.db.Get([]byte(key), nil)
	if err == leveldb.ErrNotFound {
		return "", fmt.Errorf("key '%s' not found", key)
	}
	if err != nil {
		return "", fmt.Errorf("failed to get value: %w", err)
	}
	return string(value), nil
}

func (m *Manager) Put(key, value string) error {
	err := m.db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		return fmt.Errorf("failed to put value: %w", err)
	}
	return nil
}

func (m *Manager) Delete(key string) error {
	exists, err := m.db.Has([]byte(key), nil)
	if err != nil {
		return fmt.Errorf("failed to check key existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("key '%s' not found", key)
	}

	err = m.db.Delete([]byte(key), nil)
	if err != nil {
		return fmt.Errorf("failed to delete key: %w", err)
	}
	return nil
}

func (m *Manager) List(prefix string) ([]KeyValue, error) {
	var iter iterator.Iterator
	if prefix != "" {
		iter = m.db.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	} else {
		iter = m.db.NewIterator(nil, nil)
	}
	defer iter.Release()

	var results []KeyValue
	for iter.Next() {
		results = append(results, KeyValue{
			Key:   string(iter.Key()),
			Value: string(iter.Value()),
		})
	}

	if err := iter.Error(); err != nil {
		return nil, fmt.Errorf("failed to iterate: %w", err)
	}

	return results, nil
}

func (m *Manager) Exists(key string) (bool, error) {
	exists, err := m.db.Has([]byte(key), nil)
	if err != nil {
		return false, fmt.Errorf("failed to check key existence: %w", err)
	}
	return exists, nil
}
