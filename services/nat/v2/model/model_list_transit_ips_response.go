package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransitIpsResponse Response Object
type ListTransitIpsResponse struct {

	// 查询中转IP列表的响应体。
	TransitIps *[]TransitIp `json:"transit_ips,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTransitIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpsResponse struct{}"
	}

	return strings.Join([]string{"ListTransitIpsResponse", string(data)}, " ")
}
