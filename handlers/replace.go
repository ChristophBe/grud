package handlers

import (
	"context"
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

func (c crudHandlersImpl) Replace(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	r = r.WithContext(ctx)

	model, err := c.service.GetOne(r)
	var dto types.Dto
	if dto, err = c.service.ParseDtoFromRequest(r); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = dto.IsValid(ctx, false); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if model, err = dto.AssignToModel(ctx, c.service.CreateEmptyModel(ctx)); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if model, err = model.Update(ctx); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = c.responseWriter(model, http.StatusAccepted, w, r); err != nil {
		c.errorWriter(err, w, r)
	}
}
