package eventapp

import (
	"context"
	"fmt"
	"log"

	"eventappliance/cbtypes"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/functions/metadata"
)

var bqClient bigquery.Client

func FsToBq(ctx context.Context, e cbtypes.FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}

	log.Printf("%v", meta)
	return nil
}
