package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRdsDbFaultPolicyRequest Request Object
type SetRdsDbFaultPolicyRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	Body *RdsDbFaultPolicyReq `json:"body,omitempty"`
}

func (o SetRdsDbFaultPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRdsDbFaultPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetRdsDbFaultPolicyRequest", string(data)}, " ")
}
