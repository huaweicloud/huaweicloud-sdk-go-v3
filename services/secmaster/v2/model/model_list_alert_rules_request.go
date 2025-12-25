package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRulesRequest Request Object
type ListAlertRulesRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Offset *int64 `json:"offset,omitempty"`

	// **参数解释：** 查询数据限制 **取值范围：** 0-1000 **默认取值：** 不涉及
	Limit *int64 `json:"limit,omitempty"`

	// 按照属性排序。
	SortKey *string `json:"sort_key,omitempty"`

	// 排序顺序，支持 `ASC` 或 `DESC`。
	SortDir *string `json:"sort_dir,omitempty"`

	// 输出表 ID
	OutputTableId *string `json:"output_table_id,omitempty"`

	// 告警规则名称
	AlertRuleName *string `json:"alert_rule_name,omitempty"`

	// 告警规则 ID
	AlertRuleId *string `json:"alert_rule_id,omitempty"`
}

func (o ListAlertRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAlertRulesRequest", string(data)}, " ")
}
