package storage

import (
	"fmt"

	"github.com/Scrackc/api-crud-go/model"
)

// Memory .
type Memory struct {
	currenId int
	Persons  map[int]model.Person
}

// NewMemory .
func NewMemory() Memory {
	persons := make(map[int]model.Person)

	return Memory{
		currenId: 0,
		Persons:  persons,
	}
}

// Create crea una nueva persona
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	m.currenId++
	m.Persons[m.currenId] = *person

	return nil
}

// Update actualiza una persona
func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIdPersonDoesNotExists)
	}
	m.Persons[ID] = *person

	return nil
}

// Delete elimina una persona
func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIdPersonDoesNotExists)
	}
	delete(m.Persons, ID)

	return nil
}

// GetByID retorna una persona por el ID
func (m *Memory) GetByID(ID int) (model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID: %d: %W", ID, model.ErrIdPersonDoesNotExists)
	}

	return person, nil
}

// GetAll rertorna todas las personas
func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}

	return result, nil
}
