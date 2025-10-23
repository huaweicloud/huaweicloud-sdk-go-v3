package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchAlarmRulesBody struct {

	// 当前状态是否启用。
	AlarmRuleEnable bool `json:"alarm_rule_enable"`

	// 告警规则id。
	AlarmRuleId int64 `json:"alarm_rule_id"`

	// 告警规则名称。
	AlarmRuleName string `json:"alarm_rule_name"`

	// 告警规则类型。 - metric：Prometheus指标 - event： 事件
	AlarmRuleType string `json:"alarm_rule_type"`

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml) 。  批量启停或批量修改单个企业项目下实例，填写企业项目id。
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o BatchAlarmRulesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAlarmRulesBody struct{}"
	}

	return strings.Join([]string{"BatchAlarmRulesBody", string(data)}, " ")
}
