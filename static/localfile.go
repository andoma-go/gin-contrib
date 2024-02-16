package static

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/andoma-go/gin"
)

type localFileSystem struct {
	http.FileSystem
	root    string
	indexes bool
}

func (l *localFileSystem) Exists(prefix string, file string) bool {
	if p := strings.TrimPrefix(file, prefix); len(p) < len(file) {
		name := path.Join(l.root, path.Clean(p))
		stats, err := os.Stat(name)
		if err != nil {
			return false
		}
		if stats.IsDir() {
			if !l.indexes {
				_, err := os.Stat(path.Join(name, "index.html"))
				if err != nil {
					return false
				}
			}
		}
		return true
	}
	return false
}

// Dummy
func (l *localFileSystem) Override(*gin.Context) bool {
	return false
}

func LocalFile(root string, indexes bool) *localFileSystem {
	return &localFileSystem{
		FileSystem: gin.Dir(root, indexes),
		root:       root,
		indexes:    indexes,
	}
}
