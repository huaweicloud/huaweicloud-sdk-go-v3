package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateAuthorizationRequest Request Object
type ValidateAuthorizationRequest struct {

	// **参数解释**：工作空间ID。获取方法请参见[[查询工作空间列表](ListWorkspace.xml)](tag:hc,hk)。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	WorkspaceId string `json:"workspace_id"`

	Body *ValidateAuthorizationRequestBody `json:"body,omitempty"`
}

func (o ValidateAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"ValidateAuthorizationRequest", string(data)}, " ")
}
