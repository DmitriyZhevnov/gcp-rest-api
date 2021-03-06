package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
)

func NewClient(ctx context.Context, projectID string) (*firestore.Client, error) {
	return firestore.NewClient(ctx, projectID)
}
