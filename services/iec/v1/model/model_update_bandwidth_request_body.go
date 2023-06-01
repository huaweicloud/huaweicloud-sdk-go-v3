package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新带宽请求体。
type UpdateBandwidthRequestBody struct {
	Bandwidth *UpdateBandwidthOption `json:"bandwidth"`
}

func (o UpdateBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthRequestBody", string(data)}, " ")
}
