package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthResponse Response Object
type ListBandwidthResponse struct {

	// 带宽列表对象
	EipBandwidths *[]BandwidthResponseBody `json:"eip_bandwidths,omitempty"`

	PageInfo *PageInfoOption `json:"page_info,omitempty"`

	// 本次请求的编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthResponse", string(data)}, " ")
}
