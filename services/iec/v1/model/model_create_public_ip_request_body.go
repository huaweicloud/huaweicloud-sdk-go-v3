package model

import (
	"encoding/json"

	"strings"
)

// 创建弹性公网IP请求体。
type CreatePublicIpRequestBody struct {
	Publicip *CreatePublicIpOption `json:"publicip"`
}

func (o CreatePublicIpRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePublicIpRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePublicIpRequestBody", string(data)}, " ")
}
