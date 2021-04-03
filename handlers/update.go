package handlers

import (
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

func (c crudHandlersImpl) Update(w http.ResponseWriter, r *http.Request) {

	dto, err := c.service.ParseDtoFromRequest(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = dto.IsValid(true); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	var model types.Model
	if model, err = c.service.GetOne(r); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	model, err = dto.AssignToModel(model)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if model, err = model.Update(); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = c.responseWriter(model, http.StatusAccepted, w, r); err != nil {
		c.errorWriter(err, w, r)
	}
}
