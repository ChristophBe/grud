package handlers

import (
	"context"
	"net/http"
)

func (c crudHandlersImpl) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	r = r.WithContext(ctx)

	dto, err := c.service.ParseDtoFromRequest(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = dto.IsValid(ctx, false); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	model, err := dto.AssignToModel(ctx, c.service.CreateEmptyModel(ctx))
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if model, err = model.Create(ctx); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = c.responseWriter(model, http.StatusAccepted, w, r); err != nil {
		c.errorWriter(err, w, r)
	}
}
