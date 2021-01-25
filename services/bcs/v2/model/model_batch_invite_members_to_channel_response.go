package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchInviteMembersToChannelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchInviteMembersToChannelResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchInviteMembersToChannelResponse struct{}"
	}

	return strings.Join([]string{"BatchInviteMembersToChannelResponse", string(data)}, " ")
}
