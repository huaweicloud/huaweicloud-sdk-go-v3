package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSessionRequest struct {
	NodeId string                    `json:"node_id"`
	Body   *DeleteSessionRequestBody `json:"body,omitempty"`
}

func (o DeleteSessionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSessionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSessionRequest", string(data)}, " ")
}
