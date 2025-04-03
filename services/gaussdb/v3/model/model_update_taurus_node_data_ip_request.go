package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaurusNodeDataIpRequest Request Object
type UpdateTaurusNodeDataIpRequest struct {

	// **参数解释**：               请求语言类型。   **约束限制**：   不涉及。   **取值范围**：  - en-us  - zh-cn   **默认值**：   en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：              内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认值**：  不涉及。
	ContentType string `json:"Content-Type"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  节点ID。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为no07，长度为36个字符。  **默认取值**：  不涉及。
	NodeId string `json:"node_id"`

	Body *ModifyNodeDataIpRequest `json:"body,omitempty"`
}

func (o UpdateTaurusNodeDataIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaurusNodeDataIpRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaurusNodeDataIpRequest", string(data)}, " ")
}
