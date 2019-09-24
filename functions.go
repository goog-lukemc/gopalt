package gcf

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/goog-lukemc/gopalt/types"

	"strings"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/functions/metadata"
	"google.golang.org/api/googleapi"
	//"cloud.google.com/go/functions/metadata"
)

var bqClient *bigquery.Client
var fsClient *firestore.Client
var dataset string

func init() {
	var err error
	bqClient, err = bigquery.NewClient(context.Background(), os.Getenv("BQ_PROJECT"))
	if err != nil {
		panic(err)
	}
	fsClient, err = firestore.NewClient(context.Background(), os.Getenv("GCP_PROJECT"))
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
		return unWrapLogger(bqClient.Dataset(dataset).Table("session").Uploader().Put(context.Background(), &s))
	}

	if c, ok := e.Value.Fields.(types.Conversation); ok {
		return unWrapLogger(bqClient.Dataset(dataset).Table("conversation").Uploader().Put(context.Background(), &c))
	}

	return nil
}

func unWrapLogger(err error) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(bigquery.PutMultiError); ok {
		for _, d := range e {
			log.Printf("%+v\n", d)
		}
	}

	if e, ok := err.(*googleapi.Error); ok {
		log.Printf("%+v\n", e)
	}

	return err
}

func AtMention(ctx context.Context, e types.FirestoreEvent) error {

	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}

	if c, ok := e.Value.Fields.(types.Conversation); ok {
		if strings.Contains("@mark", strings.ToLower(c.Phrase)) {
			log.Println(meta.Resource.Name)
		}
	}

	return nil

}
