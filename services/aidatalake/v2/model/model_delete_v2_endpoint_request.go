package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2EndpointRequest Request Object
type DeleteV2EndpointRequest struct {

	// **参数解释**：**参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：端点ID，可通过[查询端点列表](ListV2Endpoints.xml)获取端点ID。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	EndpointId string `json:"endpoint_id"`

	// **参数解释**：强制删除。 **约束限制**：当配置为true时，不管任何情况强制删除。 **取值范围**：   - true：强制删除。   - false：非强制删除。 **默认取值**：false。
	Force *bool `json:"force,omitempty"`
}

func (o DeleteV2EndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2EndpointRequest struct{}"
	}

	return strings.Join([]string{"DeleteV2EndpointRequest", string(data)}, " ")
}
