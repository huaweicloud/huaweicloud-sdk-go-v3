package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTemplateGroupRequest struct {
	GroupId string `json:"group_id"`
}

func (o DeleteTemplateGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateGroupRequest", string(data)}, " ")
}
