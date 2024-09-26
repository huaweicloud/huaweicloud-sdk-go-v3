package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportBindingConnectionBandwidthsResponse Response Object
type ListSupportBindingConnectionBandwidthsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 全域互联带宽列表响应体。
	GlobalconnectionBandwidths []GlobalConnectionBandwidth `json:"globalconnection_bandwidths"`
	HttpStatusCode             int                         `json:"-"`
}

func (o ListSupportBindingConnectionBandwidthsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportBindingConnectionBandwidthsResponse struct{}"
	}

	return strings.Join([]string{"ListSupportBindingConnectionBandwidthsResponse", string(data)}, " ")
}
