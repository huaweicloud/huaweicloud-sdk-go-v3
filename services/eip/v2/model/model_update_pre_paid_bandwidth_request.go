package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePrePaidBandwidthRequest struct {

	// 带宽唯一标识。通过弹性公网IP详情获取，且此弹性公网IP是包周期的。
	BandwidthId string `json:"bandwidth_id" xml:"bandwidth_id"`

	Body *UpdatePrePaidBandwidthRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdatePrePaidBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrePaidBandwidthRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrePaidBandwidthRequest", string(data)}, " ")
}
