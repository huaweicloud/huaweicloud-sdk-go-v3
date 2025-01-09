package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplySubnetBandwidthRequest Request Object
type ApplySubnetBandwidthRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *ApplySubnetBandwidthReq `json:"body,omitempty"`
}

func (o ApplySubnetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplySubnetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"ApplySubnetBandwidthRequest", string(data)}, " ")
}
