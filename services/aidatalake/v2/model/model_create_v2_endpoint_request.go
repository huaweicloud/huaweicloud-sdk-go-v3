package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateV2EndpointRequest Request Object
type CreateV2EndpointRequest struct {

	// **参数解释**：**参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	Body *CreateV2EndpointRequestBody `json:"body,omitempty"`
}

func (o CreateV2EndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateV2EndpointRequest struct{}"
	}

	return strings.Join([]string{"CreateV2EndpointRequest", string(data)}, " ")
}
