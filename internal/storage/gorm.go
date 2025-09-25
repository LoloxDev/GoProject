package storage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"GoProject/internal/contacts"

	sqlite "github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	path string
	db   *gorm.DB
}

func NewGORMStore(path string) *GORMStore {
	return &GORMStore{path: path}
}

func (s *GORMStore) Init() error {
	if s.path == "" {
		return errors.New("un chemin de base de données est requis pour GORM")
	}

	if dir := filepath.Dir(s.path); dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	dsn := fmt.Sprintf("file:%s?_pragma=foreign_keys(ON)&_pragma=busy_timeout(5000)&_journal_mode=WAL", s.path)

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&contacts.Contact{}); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *GORMStore) Create(contact *contacts.Contact) error {
	return s.db.Create(contact).Error
}

func (s *GORMStore) List() ([]contacts.Contact, error) {
	var results []contacts.Contact
	if err := s.db.Order("id asc").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (s *GORMStore) Get(id uint) (*contacts.Contact, error) {
	var contact contacts.Contact
	result := s.db.First(&contact, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, contacts.ErrNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &contact, nil
}

func (s *GORMStore) Update(contact *contacts.Contact) error {
	return s.db.Save(contact).Error
}

func (s *GORMStore) Delete(id uint) error {
	result := s.db.Delete(&contacts.Contact{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return contacts.ErrNotFound
	}
	return nil
}
