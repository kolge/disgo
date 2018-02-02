package party

import (
	"context"
	"github.com/dispatchlabs/disgover"
	"fmt"
)

type Party struct {
	members map[string]string
}

func NewParty() *Party {
	return &Party{}
}

func (party *Party) GetVersion(ctx context.Context, in *Empty) (*Version, error) {
	return &Version{"1.0.0"}, nil
}

func Join(contact *disgover.Contact) (string) {
	return fmt.Sprintf("welcome to the party %s \n", contact.GetId())
}

