package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInterRegionBandwidthRequest Request Object
type CreateInterRegionBandwidthRequest struct {
	Body *CreateInterRegionBandwidthRequestBody `json:"body,omitempty"`
}

func (o CreateInterRegionBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInterRegionBandwidthRequest struct{}"
	}

	return strings.Join([]string{"CreateInterRegionBandwidthRequest", string(data)}, " ")
}
