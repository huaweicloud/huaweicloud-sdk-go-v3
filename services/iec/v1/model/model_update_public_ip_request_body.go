package model

import (
	"encoding/json"

	"strings"
)

// 更新弹性公网IP请求数据
type UpdatePublicIpRequestBody struct {
	Publicip *UpdatePublicIpOption `json:"publicip"`
}

func (o UpdatePublicIpRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePublicIpRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePublicIpRequestBody", string(data)}, " ")
}
