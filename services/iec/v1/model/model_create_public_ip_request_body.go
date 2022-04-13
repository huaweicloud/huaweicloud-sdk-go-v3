package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建弹性公网IP请求体。
type CreatePublicIpRequestBody struct {
	Publicip *CreatePublicIpOption `json:"publicip"`
}

func (o CreatePublicIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicIpRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePublicIpRequestBody", string(data)}, " ")
}
