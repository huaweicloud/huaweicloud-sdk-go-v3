package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetBackupRetainPolicyRequestBody **参数解释**  设置备份保留策略请求体  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
type SetBackupRetainPolicyRequestBody struct {

	// **参数解释**  自动备份保留策略。NONE不保留，LAST保留最后一个，ALL全部保留。  **约束限制**  不涉及       **取值范围**  NONE、LAST、ALL  **默认取值**  不涉及。
	Auto string `json:"auto"`

	// **参数解释**  手动备份保留策略。NONE不保留，LAST保留最后一个，ALL全部保留。  **约束限制**  不涉及      **取值范围**  NONE、LAST、ALL  **默认取值**  不涉及。
	Manual string `json:"manual"`

	// **参数解释**  实例ID列表，实例ID是实例的唯一标识。  **约束限制**  不涉及。  **取值范围**  实例ID只能由英文字母、数字组成，长度为36个字符。  **默认取值**  不涉及。
	Instanceids []string `json:"instanceids"`
}

func (o SetBackupRetainPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBackupRetainPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetBackupRetainPolicyRequestBody", string(data)}, " ")
}
