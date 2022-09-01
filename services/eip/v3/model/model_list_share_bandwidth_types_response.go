package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListShareBandwidthTypesResponse struct {

	// 功能说明：共享带宽类型对象
	ShareBandwidthTypes *[]ShareBandwidthTypeShowResp `json:"share_bandwidth_types,omitempty" xml:"share_bandwidth_types"`

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	PageInfo       *PageInfoOption `json:"page_info,omitempty" xml:"page_info"`
	HttpStatusCode int             `json:"-"`
}

func (o ListShareBandwidthTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareBandwidthTypesResponse struct{}"
	}

	return strings.Join([]string{"ListShareBandwidthTypesResponse", string(data)}, " ")
}
