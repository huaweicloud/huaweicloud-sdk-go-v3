package model

import (
	"encoding/json"

	"strings"
)

type ResetUserPasswrodReq struct {
	// 用户新密码。

	NewPassword *string `json:"new_password,omitempty"`
}

func (o ResetUserPasswrodReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetUserPasswrodReq struct{}"
	}

	return strings.Join([]string{"ResetUserPasswrodReq", string(data)}, " ")
}
