package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBandwidthRequestBody 更新带宽对象的请求体(name,size必须有一个参数)
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
