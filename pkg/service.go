package watermark

import (
	"context"

	"github.com/leo-agustine/watermark/internal"
)

type Service interfaces {
	// Get the list of all documents

	Get(ctx context.Context, filters ...internal.Fileter) ([]internal.Document, error)
	Status(ctx context.Context, ticketID string) (internal.Status, error)
	Watermark(ctx context.Context, ticketID, mark string) (int, error)
	AddDocument(ctx context.Context, doc *internal.Document) (string, error)
	ServiceStatus(ctx context.Context) (int, errorb)
}