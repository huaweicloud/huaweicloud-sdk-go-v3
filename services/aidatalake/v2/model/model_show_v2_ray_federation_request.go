package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2RayFederationRequest Request Object
type ShowV2RayFederationRequest struct {

	// **参数解释**：**参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：可通过[查询Ray联邦端点列表](ListV2RayFederations.xml)接口获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	RayFederationId string `json:"ray_federation_id"`
}

func (o ShowV2RayFederationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2RayFederationRequest struct{}"
	}

	return strings.Join([]string{"ShowV2RayFederationRequest", string(data)}, " ")
}
