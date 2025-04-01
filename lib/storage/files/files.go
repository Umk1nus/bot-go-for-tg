package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/Umk1nus/bot-go-for-tg/lib"
	"github.com/Umk1nus/bot-go-for-tg/lib/storage"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Storage struct {
	basePath string
}

const (
	defaultPerm = 0774
)

var ErrNoSavedPages = errors.New("Нет сохраненных страниц")

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = lib.ErrorValidate("Не удалось сохранить", err) }()

	filePath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}

	fileName, err := createFileName(page)
	if err != nil {
		return err
	}

	filePath = filepath.Join(filePath, fileName)

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = lib.ErrorValidate("Не удалось получить", err) }()
	filePath := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(filePath)

	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, ErrNoSavedPages
	}

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(filePath, file.Name()))
}

func (s Storage) Remove(p *storage.Page) error {
	fileName, err := createFileName(p)

	if err != nil {
		return lib.ErrorValidate("Не удалось удалить файл", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("Не удалось удалить файл %s", path)
		return lib.ErrorValidate(msg, err)
	}

	return nil
}

func (s Storage) IsExists(p *storage.Page) (bool, error) {
	fileName, err := createFileName(p)
	if err != nil {
		return false, lib.ErrorValidate("Не удалось проверить существует ли файл", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("Не удалось проверить существует ли файл %s", path)
		return false, lib.ErrorValidate(msg, err)
	}

	return true, nil
}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, lib.ErrorValidate("Не удалось декодировать", err)
	}

	defer func() { _ = f.Close() }()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, lib.ErrorValidate("Не удалось декодировать", err)
	}

	return &p, nil
}

func createFileName(p *storage.Page) (string, error) {
	return p.Hash()
}
