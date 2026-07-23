package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenerateTokenResultDto struct {

	// **参数解释：**  认证Token，用于后续对结构化文档的访问鉴权。  **取值范围：**  不涉及。
	Token *string `json:"token,omitempty"`

	// **参数解释：**  用户ID，标识获取Token的IAM用户。  **取值范围：**  不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释：**  用户名，标识获取Token的IAM用户名称。  **取值范围：**  不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释：**  应用ID，用于指定实例所属的应用。  **取值范围：**  不涉及。
	AppId *string `json:"app_id,omitempty"`
}

func (o GenerateTokenResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateTokenResultDto struct{}"
	}

	return strings.Join([]string{"GenerateTokenResultDto", string(data)}, " ")
}
