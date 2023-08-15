package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBandwidthDetailResponse Response Object
type ShowBandwidthDetailResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 带宽信息
	BandWidths     *[]Bandwidth `json:"band_widths,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowBandwidthDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBandwidthDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowBandwidthDetailResponse", string(data)}, " ")
}
