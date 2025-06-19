package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlJobDefendRulesRequest Request Object
type ListSqlJobDefendRulesRequest struct {

	// 队列名称。
	QueueName *string `json:"queue_name,omitempty"`

	// 规则名称。
	RuleName *string `json:"rule_name,omitempty"`

	// 分页偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 分页个数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSqlJobDefendRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlJobDefendRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlJobDefendRulesRequest", string(data)}, " ")
}
