package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuraStatementOperatorMetricsRequest Request Object
type ShowAuraStatementOperatorMetricsRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：会话ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	SessionId string `json:"session_id"`

	// **参数解释**：statement id。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、下划线。 **默认取值**：不涉及。
	StatementId string `json:"statement_id"`
}

func (o ShowAuraStatementOperatorMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuraStatementOperatorMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowAuraStatementOperatorMetricsRequest", string(data)}, " ")
}
