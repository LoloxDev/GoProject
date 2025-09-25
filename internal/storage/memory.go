package storage

import (
	"sync"

	"GoProject/internal/contacts"
)

type MemoryStore struct {
	mu       sync.RWMutex
	contacts map[uint]contacts.Contact
	nextID   uint
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[uint]contacts.Contact),
		nextID:   1,
	}
}

func (s *MemoryStore) Init() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.contacts == nil {
		s.contacts = make(map[uint]contacts.Contact)
	}
	if s.nextID == 0 {
		s.nextID = 1
	}
	return nil
}

func (s *MemoryStore) Create(contact *contacts.Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	contact.ID = s.nextID
	s.nextID++
	s.contacts[contact.ID] = *contact
	return nil
}

func (s *MemoryStore) List() ([]contacts.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	results := make([]contacts.Contact, 0, len(s.contacts))
	for _, contact := range s.contacts {
		results = append(results, contact)
	}
	return results, nil
}

func (s *MemoryStore) Get(id uint) (*contacts.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	contact, ok := s.contacts[id]
	if !ok {
		return nil, contacts.ErrNotFound
	}

	cpy := contact
	return &cpy, nil
}

func (s *MemoryStore) Update(contact *contacts.Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.contacts[contact.ID]; !ok {
		return contacts.ErrNotFound
	}
	s.contacts[contact.ID] = *contact
	return nil
}

func (s *MemoryStore) Delete(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.contacts[id]; !ok {
		return contacts.ErrNotFound
	}
	delete(s.contacts, id)
	return nil
}
