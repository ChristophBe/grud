package types

import "context"

// Model is the representation of one data Type that is managed by your api.
type Model interface {
	// Create is used to persist a new instance of that model.
	// It returns the created model and an error if it fails otherwise nil.
	Create(ctx context.Context) (Model, error)

	// Update is used to persist new the values of an existing model.
	// It returns the updated model and an error if it fails otherwise nil.
	Update(ctx context.Context) (Model, error)

	// Delete deletes it self from the persisted records.
	// It returns an error if it fails otherwise it returns nil.
	Delete(ctx context.Context) error
}
