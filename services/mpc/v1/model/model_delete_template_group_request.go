package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTemplateGroupRequest struct {
	// 模板组id

	GroupId string `json:"group_id"`
}

func (o DeleteTemplateGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateGroupRequest", string(data)}, " ")
}
