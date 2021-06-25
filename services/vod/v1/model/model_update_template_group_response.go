package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateTemplateGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTemplateGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTemplateGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateTemplateGroupResponse", string(data)}, " ")
}
