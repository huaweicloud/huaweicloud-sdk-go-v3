package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetTokenDto struct {

	// **参数解释**：  认证类型。  **约束限制**：  不涉及。  **取值范围**：  - read: 只读 - write: 读写  **默认取值**：  不涉及。
	AuthType string `json:"auth_type"`

	// **参数解释**：  文档ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DocId string `json:"doc_id"`
}

func (o GetTokenDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTokenDto struct{}"
	}

	return strings.Join([]string{"GetTokenDto", string(data)}, " ")
}
