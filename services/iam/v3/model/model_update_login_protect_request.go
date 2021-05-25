package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateLoginProtectRequest struct {
	// 待修改登录保护状态信息的IAM用户ID。

	UserId string `json:"user_id"`

	Body *UpdateLoginProjectReq `json:"body,omitempty"`
}

func (o UpdateLoginProtectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLoginProtectRequest struct{}"
	}

	return strings.Join([]string{"UpdateLoginProtectRequest", string(data)}, " ")
}
