package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteTemplateGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTemplateGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTemplateGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplateGroupResponse", string(data)}, " ")
}
