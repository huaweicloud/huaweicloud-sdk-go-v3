package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyArchiveRuleRequest Request Object
type ApplyArchiveRuleRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	// 存档规则的唯一标识符。
	ArchiveRuleId string `json:"archive_rule_id"`
}

func (o ApplyArchiveRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyArchiveRuleRequest struct{}"
	}

	return strings.Join([]string{"ApplyArchiveRuleRequest", string(data)}, " ")
}
