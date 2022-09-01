package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新带宽的请求体
type UpdatePrePaidBandwidthRequestBody struct {
	Bandwidth *UpdatePrePaidBandwidthOption `json:"bandwidth" xml:"bandwidth"`

	ExtendParam *UpdatePrePaidBandwidthExtendParamOption `json:"extendParam,omitempty" xml:"extendParam"`
}

func (o UpdatePrePaidBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrePaidBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePrePaidBandwidthRequestBody", string(data)}, " ")
}
