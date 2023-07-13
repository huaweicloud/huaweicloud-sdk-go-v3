package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBandwidthRequest Request Object
type UpdateBandwidthRequest struct {

	// 带宽id。
	BandwidthId string `json:"bandwidth_id"`

	Body *UpdateBandwidthRequestBody `json:"body,omitempty"`
}

func (o UpdateBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthRequest struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthRequest", string(data)}, " ")
}
