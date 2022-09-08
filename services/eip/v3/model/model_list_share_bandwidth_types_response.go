package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListShareBandwidthTypesResponse struct {

	// 功能说明：共享带宽类型对象
	ShareBandwidthTypes *[]ShareBandwidthTypeShowResp `json:"share_bandwidth_types,omitempty"`

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfoOption `json:"page_info,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListShareBandwidthTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareBandwidthTypesResponse struct{}"
	}

	return strings.Join([]string{"ListShareBandwidthTypesResponse", string(data)}, " ")
}
