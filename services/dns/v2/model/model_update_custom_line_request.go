package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateCustomLineRequest struct {
	LineId string `json:"line_id"`

	Body *UpdateCustomsLineReq `json:"body,omitempty"`
}

func (o UpdateCustomLineRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCustomLineRequest struct{}"
	}

	return strings.Join([]string{"UpdateCustomLineRequest", string(data)}, " ")
}
