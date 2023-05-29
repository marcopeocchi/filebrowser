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
	Size        int64     `json:"size"`
	SHASum      string    `json:"shaSum"`
	ModTime     time.Time `json:"modTime"`
	IsImage     bool      `json:"isImage"`
	IsVideo     bool      `json:"isVideo"`
	IsDirectory bool      `json:"isDirectory"`
}

type Response struct {
	List           *[]DirectoryEntry `json:"list"`
	UpperlevelPath string            `json:"upperLevelPath"`
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
