package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateBandwidthRequest struct {

	// 带宽id。
	BandWidthId string `json:"band_width_id" xml:"band_width_id"`

	Body *UpdateBandwidthRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthRequest struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthRequest", string(data)}, " ")
}