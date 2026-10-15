package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFederationRequest Request Object
type UpdateFederationRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：Ray联邦端点ID，可通过查询Ray联邦端点列表接口获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	RayFederationId string `json:"ray_federation_id"`

	Body *UpdateRayFederationRequestBody `json:"body,omitempty"`
}

func (o UpdateFederationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFederationRequest struct{}"
	}

	return strings.Join([]string{"UpdateFederationRequest", string(data)}, " ")
}
