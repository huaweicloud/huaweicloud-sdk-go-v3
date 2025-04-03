package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMetaStudioOrdersResponse Response Object
type CreateMetaStudioOrdersResponse struct {

	// 订购的订单结果信息列表；
	OrderId *string `json:"order_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMetaStudioOrdersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetaStudioOrdersResponse struct{}"
	}

	return strings.Join([]string{"CreateMetaStudioOrdersResponse", string(data)}, " ")
}
