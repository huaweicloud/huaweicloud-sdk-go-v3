package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBandwidthRequestBody 更新带宽请求体。
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
