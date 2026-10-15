package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuraStatementOperatorMetrics **参数解释**：Aura Job V2 算子监控响应结构。
type AuraStatementOperatorMetrics struct {

	// **参数解释**：会话id。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**：语句id。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。
	StatementId *string `json:"statement_id,omitempty"`

	// **参数解释**：端点id。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**：算子监控信息。 **取值范围**：不涉及。
	OperatorMetrics *[]OperatorMetric `json:"operator_metrics,omitempty"`
}

func (o AuraStatementOperatorMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuraStatementOperatorMetrics struct{}"
	}

	return strings.Join([]string{"AuraStatementOperatorMetrics", string(data)}, " ")
}
