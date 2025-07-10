package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplySubnetBandwidthRequest Request Object
type ApplySubnetBandwidthRequest struct {
	Body *ApplySubnetBandwidthReq `json:"body,omitempty"`
}

func (o ApplySubnetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplySubnetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"ApplySubnetBandwidthRequest", string(data)}, " ")
}
