package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransitSubnetResponse Response Object
type ListTransitSubnetResponse struct {

	// 查询中转子网列表的响应体。详见TransitSubnet字段说明
	TransitSubnets *[]TransitSubnet `json:"transit_subnets,omitempty"`

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTransitSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitSubnetResponse struct{}"
	}

	return strings.Join([]string{"ListTransitSubnetResponse", string(data)}, " ")
}
