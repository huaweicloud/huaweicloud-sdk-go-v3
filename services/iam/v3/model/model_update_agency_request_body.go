package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateAgencyRequestBody struct {
	Agency *UpdateAgencyOption `json:"agency"`
}

func (o UpdateAgencyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAgencyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAgencyRequestBody", string(data)}, " ")
}
