package fs

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/marcopeocchi/filebrowser/internal/domain"
	"github.com/marcopeocchi/filebrowser/internal/utils"
)

type Repository struct {
	root string
}

func (r *Repository) WalkDir(ctx context.Context, subDir string) (domain.Response, error) {
	files := []domain.DirectoryEntry{}

	basePathLenght := strings.Count(r.root, "/")
	subDirPathLenght := strings.Count(subDir, "/")

	var upperLevel string

	err := filepath.WalkDir(
		filepath.Join(r.root, subDir),
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if strings.HasPrefix(d.Name(), ".") {
				return nil
			}

			if path == filepath.Join(r.root, subDir) {
				return nil
			}

			if strings.Count(path, "/") > basePathLenght+subDirPathLenght+1 {
				return fs.SkipDir
			}

			upperLevel = filepath.Dir(filepath.Dir(path))

			info, err := d.Info()
			if err != nil {
				return err
			}

			files = append(files, domain.DirectoryEntry{
				Path:        path,
				Name:        d.Name(),
				Size:        info.Size(),
				SHASum:      "", // TODO: implementare shasum checks
				IsVideo:     utils.IsVideo(d),
				IsDirectory: d.IsDir(),
				ModTime:     info.ModTime(),
			})

			return nil
		},
	)

	return domain.Response{
		List:           &files,
		UpperlevelPath: upperLevel,
	}, err
}

func (r *Repository) GetBasePathLength(context.Context) (int, error) {
	return strings.Count(r.root, "/"), nil
}
