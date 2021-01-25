package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowFileInfoRequest struct {
	Body *FilePath `json:"body,omitempty"`
}

func (o ShowFileInfoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowFileInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowFileInfoRequest", string(data)}, " ")
}
