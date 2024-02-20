package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateArchiveRuleRequest Request Object
type UpdateArchiveRuleRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	// 存档规则的唯一标识符。
	ArchiveRuleId string `json:"archive_rule_id"`

	Body *UpdateArchiveRuleReqBody `json:"body,omitempty"`
}

func (o UpdateArchiveRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateArchiveRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateArchiveRuleRequest", string(data)}, " ")
}
