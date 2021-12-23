package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新弹性公网IP请求数据
type UpdatePublicIpRequestBody struct {
	Publicip *UpdatePublicIpOption `json:"publicip"`
}

func (o UpdatePublicIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicIpRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePublicIpRequestBody", string(data)}, " ")
}
