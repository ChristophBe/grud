package handlers

import (
	"context"
	"net/http"
)

func (c crudHandlersImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	r = r.WithContext(ctx)

	model, err := c.service.GetAll(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = c.responseWriter(model, http.StatusOK, w, r); err != nil {
		c.errorWriter(err, w, r)
	}

}
