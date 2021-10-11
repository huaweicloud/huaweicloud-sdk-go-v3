package model

import (
	"encoding/json"

	"strings"
)

type VpcMemberCreate struct {
	// 后端实例列表

	Members []MemberInfo `json:"members"`
}

func (o VpcMemberCreate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VpcMemberCreate struct{}"
	}

	return strings.Join([]string{"VpcMemberCreate", string(data)}, " ")
}
