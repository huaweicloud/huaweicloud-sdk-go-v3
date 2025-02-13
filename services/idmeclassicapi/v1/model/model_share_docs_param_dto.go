package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareDocsParamDto struct {

	// **参数解释**：  结构化文档ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StructuredDocId *string `json:"structured_doc_id,omitempty"`

	// **参数解释**：  被分享用户UserId。  **约束限制**：  不涉及。  **取值范围**：  all：表示所有人。  **默认取值**：  不涉及。
	SharedUserId *string `json:"shared_user_id,omitempty"`

	// **参数解释**：  被分享用户名。  **约束限制**：  不涉及。  **取值范围**：  all：表示所有人。  **默认取值**：  不涉及。
	SharedUserName *string `json:"shared_user_name,omitempty"`

	// **参数解释**：  分享用户UserId。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ShareUserId *string `json:"share_user_id,omitempty"`

	// **参数解释**：  分享用户名。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ShareUserName *string `json:"share_user_name,omitempty"`

	// **参数解释**：  认证类型。  **约束限制**：  不涉及。  **取值范围**：  - read：只读。 - write：读写。  **默认取值**：  不涉及。
	AuthType *string `json:"auth_type,omitempty"`

	// **参数解释**：  修改人。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o ShareDocsParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareDocsParamDto struct{}"
	}

	return strings.Join([]string{"ShareDocsParamDto", string(data)}, " ")
}
