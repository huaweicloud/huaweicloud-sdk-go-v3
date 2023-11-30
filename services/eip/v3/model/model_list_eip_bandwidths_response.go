package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEipBandwidthsResponse Response Object
type ListEipBandwidthsResponse struct {

	// 带宽列表对象
	EipBandwidths *[]EipBandwidthResponseBody `json:"eip_bandwidths,omitempty"`

	PageInfo *PageInfoDict `json:"page_info,omitempty"`

	// 本次请求的编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListEipBandwidthsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEipBandwidthsResponse struct{}"
	}

	return strings.Join([]string{"ListEipBandwidthsResponse", string(data)}, " ")
}
