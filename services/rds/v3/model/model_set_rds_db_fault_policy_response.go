package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRdsDbFaultPolicyResponse Response Object
type SetRdsDbFaultPolicyResponse struct {

	// **参数解释**：  请求状态。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	State *string `json:"state,omitempty"`

	// **参数解释**：  错误信息。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Errmsg         *string `json:"errmsg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetRdsDbFaultPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRdsDbFaultPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetRdsDbFaultPolicyResponse", string(data)}, " ")
}
