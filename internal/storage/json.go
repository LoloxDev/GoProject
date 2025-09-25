package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"GoProject/internal/contacts"
)

type JSONStore struct {
	path     string
	mu       sync.RWMutex
	contacts map[uint]contacts.Contact
	nextID   uint
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{path: path}
}

func (s *JSONStore) Init() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.contacts == nil {
		s.contacts = make(map[uint]contacts.Contact)
	}

	if s.nextID == 0 {
		s.nextID = 1
	}

	if s.path == "" {
		return nil
	}

	if dir := filepath.Dir(s.path); dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	data, err := os.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var payload struct {
		Contacts []contacts.Contact `json:"contacts"`
		NextID   uint               `json:"next_id"`
	}

	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	for _, contact := range payload.Contacts {
		s.contacts[contact.ID] = contact
		if contact.ID >= s.nextID {
			s.nextID = contact.ID + 1
		}
	}

	if payload.NextID >= s.nextID {
		s.nextID = payload.NextID
	}

	return nil
}

func (s *JSONStore) Create(contact *contacts.Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	contact.ID = s.nextID
	s.nextID++
	s.contacts[contact.ID] = *contact

	return s.persist()
}

func (s *JSONStore) List() ([]contacts.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	results := make([]contacts.Contact, 0, len(s.contacts))
	for _, contact := range s.contacts {
		results = append(results, contact)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].ID < results[j].ID
	})

	return results, nil
}

func (s *JSONStore) Get(id uint) (*contacts.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	contact, ok := s.contacts[id]
	if !ok {
		return nil, contacts.ErrNotFound
	}

	cpy := contact
	return &cpy, nil
}

func (s *JSONStore) Update(contact *contacts.Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.contacts[contact.ID]; !ok {
		return contacts.ErrNotFound
	}

	s.contacts[contact.ID] = *contact
	return s.persist()
}

func (s *JSONStore) Delete(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.contacts[id]; !ok {
		return contacts.ErrNotFound
	}

	delete(s.contacts, id)
	return s.persist()
}

func (s *JSONStore) persist() error {
	if s.path == "" {
		return nil
	}

	payload := struct {
		Contacts []contacts.Contact `json:"contacts"`
		NextID   uint               `json:"next_id"`
	}{
		Contacts: make([]contacts.Contact, 0, len(s.contacts)),
		NextID:   s.nextID,
	}

	for _, contact := range s.contacts {
		payload.Contacts = append(payload.Contacts, contact)
	}

	sort.Slice(payload.Contacts, func(i, j int) bool {
		return payload.Contacts[i].ID < payload.Contacts[j].ID
	})

	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}

	if dir := filepath.Dir(s.path); dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	return os.WriteFile(s.path, data, 0o644)
}
