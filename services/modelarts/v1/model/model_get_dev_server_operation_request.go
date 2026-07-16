package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDevServerOperationRequest Request Object
type GetDevServerOperationRequest struct {

	// **参数解释**：Lite Server实例ID。 **约束限制**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id string `json:"id"`

	// **参数解释**：Operation ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OperationId string `json:"operation_id"`
}

func (o GetDevServerOperationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDevServerOperationRequest struct{}"
	}

	return strings.Join([]string{"GetDevServerOperationRequest", string(data)}, " ")
}
