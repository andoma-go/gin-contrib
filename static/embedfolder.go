package static

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/andoma-go/gin"
)

type embedFileSystem struct {
	http.FileSystem
	overrides []Override
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func (e embedFileSystem) Override(c *gin.Context) bool {
	if len(e.overrides) > 0 {
		fn := e.overrides[0]
		return fn(c)
	}
	return false
}

func EmbedFolder(fsEmbed embed.FS, reqPath string, overrides ...Override) (ServeFileSystem, error) {
	targetPath := strings.TrimSpace(reqPath)
	if targetPath == "" {
		return embedFileSystem{
			FileSystem: http.FS(fsEmbed),
			overrides:  overrides,
		}, nil
	}

	fsys, _ := fs.Sub(fsEmbed, targetPath)
	_, err := fsEmbed.Open(targetPath)
	if err != nil {
		return nil, err
	}

	return embedFileSystem{
		FileSystem: http.FS(fsys),
		overrides:  overrides,
	}, nil
}

func ServeEmbed(reqPath string, fsEmbed embed.FS) gin.HandlerFunc {
	embedFS, err := EmbedFolder(fsEmbed, reqPath)
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "initialization of embed folder failed",
				"error":   err.Error(),
			})
		}
	}
	return gin.WrapH(http.FileServer(embedFS))
}
