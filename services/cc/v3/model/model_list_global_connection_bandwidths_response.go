package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalConnectionBandwidthsResponse Response Object
type ListGlobalConnectionBandwidthsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 全域互联带宽列表响应体。
	GlobalconnectionBandwidths []GlobalConnectionBandwidth `json:"globalconnection_bandwidths"`
	HttpStatusCode             int                         `json:"-"`
}

func (o ListGlobalConnectionBandwidthsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConnectionBandwidthsResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalConnectionBandwidthsResponse", string(data)}, " ")
}
