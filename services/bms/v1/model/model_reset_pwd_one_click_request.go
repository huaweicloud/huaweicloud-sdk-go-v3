package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResetPwdOneClickRequest struct {
	// 裸金属服务器ID。可以从裸金属服务器控制台查询，或者通过调用7.3.4-查询裸金属服务器列表（OpenStack原生）API获取。

	ServerId string `json:"server_id"`

	Body *ResetPasswordBody `json:"body,omitempty"`
}

func (o ResetPwdOneClickRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetPwdOneClickRequest struct{}"
	}

	return strings.Join([]string{"ResetPwdOneClickRequest", string(data)}, " ")
}
