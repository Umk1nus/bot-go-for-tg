package files

import (
	"github.com/Umk1nus/bot-go-for-tg/lib"
	"github.com/Umk1nus/bot-go-for-tg/lib/storage"
	"os"
	"path/filepath"
)

type Storage struct {
	basePath string
}

const (
	defaultPerm = 0774
)

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = lib.ErrorValidate("Не удалось сохранить", err) }()

	filePath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}
	return
}
