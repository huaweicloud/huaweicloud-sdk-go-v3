package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthTypesResponse Response Object
type ListBandwidthTypesResponse struct {

	// 共享带宽类型列表对象。
	ShareBandwidthTypes *[]BandwidthTypeOption `json:"share_bandwidth_types,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBandwidthTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthTypesResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthTypesResponse", string(data)}, " ")
}
