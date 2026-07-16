package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrderRequestBody 创建订单的数据模型。
type CreateOrderRequestBody struct {

	// **参数解释**：订单类型。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ActionType *string `json:"actionType,omitempty"`
}

func (o CreateOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderRequestBody struct{}"
	}

	return strings.Join([]string{"CreateOrderRequestBody", string(data)}, " ")
}
