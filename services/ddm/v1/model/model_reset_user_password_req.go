package model

import (
	"encoding/json"

	"strings"
)

// ResetUserPasswordReq。
type ResetUserPasswordReq struct {
	// 重置后的新密码。

	Password string `json:"password"`
}

func (o ResetUserPasswordReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetUserPasswordReq struct{}"
	}

	return strings.Join([]string{"ResetUserPasswordReq", string(data)}, " ")
}
