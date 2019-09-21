package gcf

import (
	"context"
	"strconv"
	"testing"
	"time"

	"cloud.google.com/go/functions/metadata"
	"github.com/goog-lukemc/gopalt/types"
)

func TestEventToBq(t *testing.T) {
	meta := metadata.Metadata{
		EventID:   "abc1234",
		Timestamp: time.Now(),
		EventType: "bob",
		Resource: &metadata.Resource{
			Service: "service",
			Name:    "name",
			Type:    "resource type",
		},
	}

	sess := types.FirestoreEvent{
		Value: types.FirestoreValue{
			Fields: NewSession(),
		},
	}

	ctx := metadata.NewContext(context.Background(), &meta)
	err := FsToBq(ctx, sess)
	if err != nil {
		t.Fatalf("%+v", err.Error())
	}

}

func NewSession() types.Session {
	return types.Session{
		UserId:     strconv.FormatInt(time.Now().UnixNano(), 10),
		SessionId:  time.Now().UnixNano(),
		DialogueId: time.Now().UnixNano(),
	}
}

func NewConversation(s *types.Session) *types.Conversation {
	return &types.Conversation{
		DialogueId: s.DialogueId,
		Phrase:     strconv.FormatInt(time.Now().UnixNano(), 10),
		Response:   strconv.FormatInt(time.Now().UnixNano()+1, 10),
	}
}
