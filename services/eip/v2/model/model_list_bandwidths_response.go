package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthsResponse Response Object
type ListBandwidthsResponse struct {

	// 带宽列表对象
	Bandwidths     *[]BandwidthResp `json:"bandwidths,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListBandwidthsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthsResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthsResponse", string(data)}, " ")
}
