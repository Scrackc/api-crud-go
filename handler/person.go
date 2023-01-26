package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Scrackc/api-crud-go/model"
)

// person
type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "Ocurrio un error al crear la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Persona creada", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un numúmero entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(id, &data)
	if err != nil {
		response := newResponse(Error, "Error al actualizar a la pesona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "Persona actualizada", nil)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	resp, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Error al obtener las personas", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", resp)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un numúmero entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(id)
	if errors.Is(err, model.ErrIdPersonDoesNotExists) {
		response := newResponse(Error, "Persona no existente", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Error al eliminar la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "Persona eliminada", nil)
	responseJSON(w, http.StatusOK, response)

}

func (p *person) getByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		respose := newResponse(Error, "El id debe ser de tipo entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, respose)
		return
	}

	person, err := p.storage.GetByID(id)
	if errors.Is(err, model.ErrIdPersonDoesNotExists) {
		response := newResponse(Error, "Persona no existente", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Error al obtener a la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", person)
	responseJSON(w, http.StatusOK, response)

}
