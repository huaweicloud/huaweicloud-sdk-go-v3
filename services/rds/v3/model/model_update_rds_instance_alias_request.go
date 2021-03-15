package model

import (
	"encoding/json"

	"strings"
)

type UpdateRdsInstanceAliasRequest struct {
	// 备注名称长度可在0~64个字符之间，由字母、数字、汉字，英文句号，下划线、中划线组成。
	Alias *string `json:"alias,omitempty"`
}

func (o UpdateRdsInstanceAliasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRdsInstanceAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdateRdsInstanceAliasRequest", string(data)}, " ")
}
