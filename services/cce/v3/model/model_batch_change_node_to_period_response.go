package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeNodeToPeriodResponse Response Object
type BatchChangeNodeToPeriodResponse struct {

	// **参数解释**： 提交订单成功后返回的订单ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	OrderId        *string `json:"orderId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchChangeNodeToPeriodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeNodeToPeriodResponse struct{}"
	}

	return strings.Join([]string{"BatchChangeNodeToPeriodResponse", string(data)}, " ")
}
