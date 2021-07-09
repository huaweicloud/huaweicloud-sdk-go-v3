package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateUsersReq struct {
	// DDM实例帐号相关信息的集合。

	Users []CreateUsersInfo `json:"users"`
}

func (o CreateUsersReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateUsersReq struct{}"
	}

	return strings.Join([]string{"CreateUsersReq", string(data)}, " ")
}
