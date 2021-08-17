package model

import (
	"encoding/json"

	"strings"
)

type AssignedUserInfo struct {
	// id信息

	Id *string `json:"id,omitempty"`
	// 名称信息

	Name *string `json:"name,omitempty"`
}

func (o AssignedUserInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssignedUserInfo struct{}"
	}

	return strings.Join([]string{"AssignedUserInfo", string(data)}, " ")
}
