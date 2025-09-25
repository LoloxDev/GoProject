package storage

import (
	"fmt"

	"GoProject/internal/config"
	"GoProject/internal/contacts"
)

type Storer interface {
	Init() error
	Create(*contacts.Contact) error
	List() ([]contacts.Contact, error)
	Get(uint) (*contacts.Contact, error)
	Update(*contacts.Contact) error
	Delete(uint) error
}

func NewFromConfig(cfg config.AppConfig) (Storer, error) {
	switch cfg.Storage.Type {
	case "gorm":
		if cfg.Storage.Gorm.Path == "" {
			return nil, fmt.Errorf("storage.gorm.path manquant")
		}
		return NewGORMStore(cfg.Storage.Gorm.Path), nil
	case "json":
		if cfg.Storage.JSON.Path == "" {
			return nil, fmt.Errorf("storage.json.path manquant")
		}
		return NewJSONStore(cfg.Storage.JSON.Path), nil
	case "memory":
		return NewMemoryStore(), nil
	default:
		return nil, fmt.Errorf("type de storage inconnu: %s", cfg.Storage.Type)
	}
}
