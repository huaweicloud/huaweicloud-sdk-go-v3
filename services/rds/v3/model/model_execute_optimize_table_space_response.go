package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteOptimizeTableSpaceResponse Response Object
type ExecuteOptimizeTableSpaceResponse struct {

	// **参数解释**：  结果。  **约束限制**：  不涉及。  **取值范围**：  -successful  **默认取值**：  不涉及。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteOptimizeTableSpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteOptimizeTableSpaceResponse struct{}"
	}

	return strings.Join([]string{"ExecuteOptimizeTableSpaceResponse", string(data)}, " ")
}
