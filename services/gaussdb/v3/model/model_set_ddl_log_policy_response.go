package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetDdlLogPolicyResponse Response Object
type SetDdlLogPolicyResponse struct {

	// **参数解释**：  任务流ID。  **取值范围**：  不涉及。
	WorkflowId     *string `json:"workflow_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetDdlLogPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDdlLogPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetDdlLogPolicyResponse", string(data)}, " ")
}
