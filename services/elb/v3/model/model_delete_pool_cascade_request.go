package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePoolCascadeRequest Request Object
type DeletePoolCascadeRequest struct {

	// **参数解释**：后端服务器组ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	PoolId string `json:"pool_id"`
}

func (o DeletePoolCascadeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePoolCascadeRequest struct{}"
	}

	return strings.Join([]string{"DeletePoolCascadeRequest", string(data)}, " ")
}
