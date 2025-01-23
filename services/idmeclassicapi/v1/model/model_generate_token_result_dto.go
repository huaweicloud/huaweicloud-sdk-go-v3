package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenerateTokenResultDto struct {

	// **参数解释**：  认证token。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Token *string `json:"token,omitempty"`

	// **参数解释**：  用户id。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**：  用户名。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**：  应用id。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	AppId *string `json:"app_id,omitempty"`
}

func (o GenerateTokenResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateTokenResultDto struct{}"
	}

	return strings.Join([]string{"GenerateTokenResultDto", string(data)}, " ")
}
