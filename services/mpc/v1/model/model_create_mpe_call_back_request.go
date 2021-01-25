package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMpeCallBackRequest struct {
	Body *MpeCallBackReq `json:"body,omitempty"`
}

func (o CreateMpeCallBackRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMpeCallBackRequest struct{}"
	}

	return strings.Join([]string{"CreateMpeCallBackRequest", string(data)}, " ")
}
