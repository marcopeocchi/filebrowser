package domain

import (
	"context"
	"net/http"
	"time"
)

const (
	OrderByName = iota
	OrderByTime
)

type DirectoryEntry struct {
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	UpperLevel  string    `json:"upperLevel"`
	Size        int64     `json:"size"`
	SHASum      string    `json:"shaSum"`
	ModTime     time.Time `json:"modTime"`
	IsImage     bool      `json:"isImage"`
	IsVideo     bool      `json:"isVideo"`
	IsDirectory bool      `json:"isDirectory"`
}

type Response struct {
	BasePathLenght int               `json:"basePathLength"`
	List           *[]DirectoryEntry `json:"list"`
}

type FSRepository interface {
	WalkDir(ctx context.Context, subDir string) (Response, error)
	GetBasePathLength(ctx context.Context) (int, error)
}

type FSService interface {
	WalkDir(ctx context.Context, subDir string) (Response, error)
	GetBasePathLength(ctx context.Context) (int, error)
}

type FSHandler interface {
	WalkDir() http.HandlerFunc
	OpenFile() http.HandlerFunc
	GetBasePathLength() http.HandlerFunc
}
