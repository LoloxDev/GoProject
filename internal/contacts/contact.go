package contacts

import "errors"

var ErrNotFound = errors.New("contact introuvable")

type Contact struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Store interface {
	Init() error
	Create(contact *Contact) error
	List() ([]Contact, error)
	Get(id uint) (*Contact, error)
	Update(contact *Contact) error
	Delete(id uint) error
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) Add(name, email, password string) (*Contact, error) {
	contact := &Contact{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := s.store.Create(contact); err != nil {
		return nil, err
	}

	return contact, nil
}

func (s *Service) List() ([]Contact, error) {
	return s.store.List()
}

func (s *Service) Update(id uint, name, email, password string) (*Contact, error) {
	contact, err := s.store.Get(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		contact.Name = name
	}
	if email != "" {
		contact.Email = email
	}
	if password != "" {
		contact.Password = password
	}

	if err := s.store.Update(contact); err != nil {
		return nil, err
	}

	return contact, nil
}

func (s *Service) Delete(id uint) error {
	return s.store.Delete(id)
}
