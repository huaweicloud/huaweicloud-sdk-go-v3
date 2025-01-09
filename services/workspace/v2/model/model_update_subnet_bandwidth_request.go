package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubnetBandwidthRequest Request Object
type UpdateSubnetBandwidthRequest struct {

	// 云办公带宽id。
	BandwidthId string `json:"bandwidth_id"`

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *UpdateSubnetBandwidthReq `json:"body,omitempty"`
}

func (o UpdateSubnetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubnetBandwidthRequest", string(data)}, " ")
}
