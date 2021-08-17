package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteServersRequest struct {
	Body *DeleteIds `json:"body,omitempty"`
}

func (o DeleteServersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteServersRequest struct{}"
	}

	return strings.Join([]string{"DeleteServersRequest", string(data)}, " ")
}
