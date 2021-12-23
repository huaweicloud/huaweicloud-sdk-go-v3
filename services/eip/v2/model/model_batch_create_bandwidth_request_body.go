package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量创建带宽的请求体
type BatchCreateBandwidthRequestBody struct {
	Bandwidth *BatchCreateBandwidthOption `json:"bandwidth"`
}

func (o BatchCreateBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateBandwidthRequestBody", string(data)}, " ")
}
