package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrderIdRequest Request Object
type CreateOrderIdRequest struct {

	// **参数解释**：资源池ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：订单操作类型。 **约束限制**：不涉及。 **取值范围**：可选值如下： - toPeriod：按需转包周期，默认值 **默认取值**：不涉及。
	ActionType *string `json:"actionType,omitempty"`

	// **参数解释**：工作空间ID，默认是0。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	WorkspaceId *string `json:"workspaceId,omitempty"`

	Body *CreateOrderRequestBody `json:"body,omitempty"`
}

func (o CreateOrderIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderIdRequest struct{}"
	}

	return strings.Join([]string{"CreateOrderIdRequest", string(data)}, " ")
}
