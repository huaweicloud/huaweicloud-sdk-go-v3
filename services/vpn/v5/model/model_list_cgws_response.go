package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCgwsResponse Response Object
type ListCgwsResponse struct {

	// 对端网关信息
	CustomerGateways *[]ResponseCustomerGateway `json:"customer_gateways,omitempty"`

	// 租户下对端网关总数
	TotalCount *int64 `json:"total_count,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCgwsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCgwsResponse struct{}"
	}

	return strings.Join([]string{"ListCgwsResponse", string(data)}, " ")
}
