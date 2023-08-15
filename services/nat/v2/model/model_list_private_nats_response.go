package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateNatsResponse Response Object
type ListPrivateNatsResponse struct {

	// 查询私网NAT网关实例列表的响应体。 详见PrivateNat字段说明。
	Gateways *[]PrivateNat `json:"gateways,omitempty"`

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPrivateNatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNatsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateNatsResponse", string(data)}, " ")
}
