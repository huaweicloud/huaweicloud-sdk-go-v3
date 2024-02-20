package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowArchiveRuleRequest Request Object
type ShowArchiveRuleRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	// 存档规则的唯一标识符。
	ArchiveRuleId string `json:"archive_rule_id"`
}

func (o ShowArchiveRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowArchiveRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowArchiveRuleRequest", string(data)}, " ")
}
