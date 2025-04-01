package storage

import (
	"crypto/sha1"
	"fmt"
	"github.com/Umk1nus/bot-go-for-tg/lib"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	isExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

func (p *Page) Hash() (string, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", lib.ErrorValidate("Не удалось получить хэш", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", lib.ErrorValidate("Не удалось получить хэш", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
