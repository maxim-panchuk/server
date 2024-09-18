package bet

import "github.com/google/uuid"

type Event struct {
	ID *uuid.UUID
}

type Bet struct {
	ID          *uuid.UUID
	Title       string
	Description string
}
