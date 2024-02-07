package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInternetBandwidthRequest Request Object
type CreateInternetBandwidthRequest struct {
	Body *CreateInternetBandwidthRequestBody `json:"body,omitempty"`
}

func (o CreateInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"CreateInternetBandwidthRequest", string(data)}, " ")
}
