package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQuotasOrderResponse Response Object
type CreateQuotasOrderResponse struct {

	// **参数解释**： 订单ID **约束限制**: 不涉及 **取值范围**： 字符长度1-256 **默认取值**: 不涉及
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateQuotasOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQuotasOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateQuotasOrderResponse", string(data)}, " ")
}
