package model

import (
	"encoding/json"

	"strings"
)

//
type MemberRef struct {
	// 后端服务器ID。
	Id string `json:"id"`
}

func (o MemberRef) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MemberRef struct{}"
	}

	return strings.Join([]string{"MemberRef", string(data)}, " ")
}
