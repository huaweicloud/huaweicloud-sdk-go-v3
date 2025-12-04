package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIntelligentKillSessionStatisticRequest Request Object
type ShowIntelligentKillSessionStatisticRequest struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  节点ID。  **约束限制**：  只有当实例为集群版时该参数有效。
	NodeId *string `json:"node_id,omitempty"`
}

func (o ShowIntelligentKillSessionStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIntelligentKillSessionStatisticRequest struct{}"
	}

	return strings.Join([]string{"ShowIntelligentKillSessionStatisticRequest", string(data)}, " ")
}
