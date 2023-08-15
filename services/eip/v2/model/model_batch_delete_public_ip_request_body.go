package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePublicIpRequestBody 请求参数
type BatchDeletePublicIpRequestBody struct {

	// 弹性公网ip的id列表。
	PublicipIds []string `json:"publicip_ids"`
}

func (o BatchDeletePublicIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicIpRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicIpRequestBody", string(data)}, " ")
}
