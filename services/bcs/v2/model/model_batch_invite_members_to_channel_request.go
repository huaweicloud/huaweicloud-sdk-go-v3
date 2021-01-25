package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchInviteMembersToChannelRequest struct {
	Body *BatchInviteMembersToChannelRequestBody `json:"body,omitempty"`
}

func (o BatchInviteMembersToChannelRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchInviteMembersToChannelRequest struct{}"
	}

	return strings.Join([]string{"BatchInviteMembersToChannelRequest", string(data)}, " ")
}
