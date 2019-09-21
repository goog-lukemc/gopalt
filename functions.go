package gcf

import (
	"context"
	"os"

	"github.com/goog-lukemc/gopalt/types"

	"cloud.google.com/go/bigquery"
	//"cloud.google.com/go/functions/metadata"
)

var bqClient *bigquery.Client
var dataset string

func init() {
	var err error
	bqClient, err = bigquery.NewClient(context.Background(), os.Getenv("BQ_PROJECT"))
	if err != nil {
		panic(err)
	}
	dataset = os.Getenv("BQ_DATASET")
}

func FsToBq(ctx context.Context, e types.FirestoreEvent) error {
	// meta, err := metadata.FromContext(ctx)
	// if err != nil {
	// 	return fmt.Errorf("metadata.FromContext: %v", err)
	// }

	if s, ok := e.Value.Fields.(types.Session); ok {
		return bqClient.Dataset(dataset).Table("session").Uploader().Put(context.Background(), &s)
	}

	if c, ok := e.Value.Fields.(types.Conversation); ok {
		return bqClient.Dataset(dataset).Table("conversation").Uploader().Put(context.Background(), &c)
	}

	return nil
}
