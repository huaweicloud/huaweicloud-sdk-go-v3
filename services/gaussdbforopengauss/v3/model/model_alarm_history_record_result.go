package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmHistoryRecordResult struct {

	// **参数解释**： 告警规则ID。 **取值范围**： 不涉及。
	AlarmId string `json:"alarm_id"`

	// **参数解释**： 告警规则名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 告警记录的状态。 **取值范围**： - ok：正常。 - alarm：告警。 - invalid：已失效。
	Status string `json:"status"`

	// **参数解释**： 告警规则类型。 **取值范围**： - EVENT.SYS：系统事件告警。 - EVENT.CUSTOM：自定义事件告警。 - DNSHealthCheck：DNS健康检查告警。 - RESOURCE_GROUP：资源分组告警。 - MULTI_INSTANCE：指定资源告警。
	AlarmType string `json:"alarm_type"`

	// **参数解释**： 告警历史的告警级别。 **取值范围**： - 1：紧急。 - 2：重要。 - 3：次要。 - 4：提示。
	Level int32 `json:"level"`

	// **参数解释**： 实例ID。 **取值范围**： 只能由英文字母、数字组成，且长度为36个字符。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	InstanceName string `json:"instance_name"`

	// **参数解释**： 告警开始时间。 **取值范围**： UNIX时间戳，单位毫秒，例如：1603131199000。
	BeginTime int64 `json:"begin_time"`

	// **参数解释**： 告警规则的最后更新时间。 **取值范围**： UNIX时间戳，单位毫秒，例如：1603131199000。
	UpdateTime int64 `json:"update_time"`
}

func (o AlarmHistoryRecordResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryRecordResult struct{}"
	}

	return strings.Join([]string{"AlarmHistoryRecordResult", string(data)}, " ")
}
