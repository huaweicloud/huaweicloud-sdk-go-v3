package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAlarmRuleRequest Request Object
type BatchUpdateAlarmRuleRequest struct {

	// 批量操作action：  - enable：批量启动Prometheus监控告警规则 - disable：批量停止Prometheus监控告警规则 - BATCH_UPDATE_ACTION_RULE：批量修改Prometheus监控告警规则的告警行动规则
	Action string `json:"action"`

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml) 。 - 批量启停或批量修改单个企业项目下实例，填写企业项目id。  - 不填 则批量启停或批量修改默认企业项目下实例，默认企业项目id为0。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`

	Body *BatchUpdateRequest `json:"body,omitempty"`
}

func (o BatchUpdateAlarmRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAlarmRuleRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateAlarmRuleRequest", string(data)}, " ")
}
