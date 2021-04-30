package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchAddPeersToChannelResponse struct {
	// 操作记录id

	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddPeersToChannelResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddPeersToChannelResponse struct{}"
	}

	return strings.Join([]string{"BatchAddPeersToChannelResponse", string(data)}, " ")
}
