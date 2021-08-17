package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAlertConfigRequest struct {
	Body *UpdateAlertConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateAlertConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAlertConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlertConfigRequest", string(data)}, " ")
}
