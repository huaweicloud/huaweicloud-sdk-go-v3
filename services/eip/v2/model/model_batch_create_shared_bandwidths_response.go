package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSharedBandwidthsResponse Response Object
type BatchCreateSharedBandwidthsResponse struct {

	// 批创的带宽对象的列表
	Bandwidths     *[]BatchBandwidthResp `json:"bandwidths,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchCreateSharedBandwidthsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSharedBandwidthsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSharedBandwidthsResponse", string(data)}, " ")
}
