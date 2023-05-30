package db

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/marcopeocchi/filebrowser/internal/domain"
)

type JsonDB struct {
	path  string
	users []domain.User
	table sync.Map
}

func NewJsonDB(path string) (*JsonDB, error) {
	j := &JsonDB{
		path: path,
	}
	err := j.loadFromFile()
	if err != nil {
		return nil, err
	}

	for _, u := range j.users {
		j.table.Store(u.Uid, u)
	}

	return j, nil
}

func (j *JsonDB) loadFromFile() error {
	fd, err := os.ReadFile(j.path)
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewReader(fd)).Decode(&j.users)
}

func (j *JsonDB) FindById(ctx context.Context, id string) (domain.User, error) {
	select {
	case <-ctx.Done():
		return domain.User{}, errors.New("context cancelled")
	default:
		val, ok := j.table.Load(id)
		if ok {
			return val.(domain.User), nil
		}
		return domain.User{}, errors.New("cannot find user with the given id")
	}
}

func (j *JsonDB) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	u := domain.User{}

	select {
	case <-ctx.Done():
		return u, errors.New("context cancelled")

	default:
		j.table.Range(func(key, value any) bool {
			if value.(domain.User).Username == username {
				u = value.(domain.User)
				return true
			}
			return false
		})

		if u.Uid != "" {
			return u, nil
		}
		return u, errors.New("cannot find user with the given id")
	}
}
