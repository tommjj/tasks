package storage

import (
	"context"
	"encoding/json"
	"os"
	"reflect"
	"sync"

	"github.com/tommjj/tasks/internal/core/domain"
	"github.com/tommjj/tasks/internal/core/ports"
)

type fileStorage struct {
	Filename string
	lock     sync.RWMutex
}

func New(filename string) ports.IStorage {
	return &fileStorage{
		Filename: filename,
	}
}

func (f *fileStorage) Load(ctx context.Context, v any) error {
	f.lock.RLock()
	defer f.lock.RUnlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return domain.ErrTypeNotPtr
	}

	kind := val.Elem().Kind()
	isSupport := kind == reflect.Slice || kind == reflect.Struct || kind == reflect.Map
	if !isSupport {
		return domain.ErrUnsupportedType
	}

	file, err := os.Open(f.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (f *fileStorage) Sync(ctx context.Context, v any) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(f.Filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
