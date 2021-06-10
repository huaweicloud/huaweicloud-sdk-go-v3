package model

import (
	"encoding/json"

	"strings"
)

type CreateMessageV2Req struct {
	Message *CreateMessageDoV2 `json:"message"`
	// 组id

	GroupId *string `json:"group_id,omitempty"`
}

func (o CreateMessageV2Req) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMessageV2Req struct{}"
	}

	return strings.Join([]string{"CreateMessageV2Req", string(data)}, " ")
}
