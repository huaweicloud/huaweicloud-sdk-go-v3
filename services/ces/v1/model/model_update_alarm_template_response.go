package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateAlarmTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmTemplateResponse", string(data)}, " ")
}
