package model

import (
	"encoding/json"

	"strings"
)

// 生成接入凭证的结构体。
type CreateAccessCodeRequestBody struct {
	// 接入凭证类型，默认为AMQP的接入凭证类型。
	Type *string `json:"type,omitempty"`
}

func (o CreateAccessCodeRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAccessCodeRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAccessCodeRequestBody", string(data)}, " ")
}
