package cbtypes

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue FirestoreValue `json:"oldValue"`
	Value    FirestoreValue `json:"value"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	Fields interface{} `json:"fields"`
}

type Session struct {
	UserId     string `firestore:"userId" bigquery:"userId"`
	SessionId  int64  `firestore:"sessionId" bigquery:"sessionId"`
	DialogueId int64  `firestore:"dialogueId" bigquery:"dialogueId"`
}

type Conversation struct {
	DialogueId int64  `firestore:"dialogueId" bigquery:"dialogueId"`
	Phrase     string `firestore:"phrase" bigquery:"phrase"`
	Response   string `firestore:"response" bigquery:"response"`
}

type Rating struct {
	SessionID  int64 `firestore:"sessionId" bigquery:"sessionId"`
	UserRating int64 `firestore:"userRating" bigquery:"userRating"`
}
