package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrderRequest Request Object
type ShowOrderRequest struct {

	// **参数解释**：订单ID。取值自订单列表返回的orderName字段。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OrderName string `json:"order_name"`
}

func (o ShowOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrderRequest struct{}"
	}

	return strings.Join([]string{"ShowOrderRequest", string(data)}, " ")
}
