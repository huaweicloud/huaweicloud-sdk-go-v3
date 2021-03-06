package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateMemberRequestBody struct {
	Member *UpdateMemberReq `json:"member"`
}

func (o UpdateMemberRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateMemberRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateMemberRequestBody", string(data)}, " ")
}
