package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupTriggerPropertiesRequestInfo1 **参数解释**: 策略执行时间规则，若开启勒索防护时开启备份功能，则该字段必选 **约束限制**: 不涉及 **取值范围**: 不涉及  **默认取值**: 不涉及
type BackupTriggerPropertiesRequestInfo1 struct {

	// **参数解释**: 调度规则。若开启勒索防护时开启备份功能，则该字段必选。 **约束限制**: 限制24条规则。 **取值范围**: 调度器的调度规则，可参照iCalendar RFC 2445规范中的事件规则，但仅支持FREQ、BYDAY、BYHOUR、BYMINUTE、INTERVAL等参数，其中FREQ仅支持WEEKLY和DAILY，BYDAY支持一周七天（MO、TU、WE、TH、FR、SA、SU），BYHOUR支持0-23小时，BYMINUTE支持0-59分钟，并且间隔不能小于一小时，一天最大24个时间点。例如，周一到周天，每天14:00调度，其规则为：'FREQ=WEEKLY;BYDAY=MO，TU，WE，TH，FR，SA，SU;BYHOUR=14;BYMINUTE=00'。每天14:00调度，其规则为'FREQ=DAILY;INTERVAL=1;BYHOUR=14;BYMINUTE=00'。  **默认取值**: 不涉及
	Pattern *[]string `json:"pattern,omitempty"`
}

func (o BackupTriggerPropertiesRequestInfo1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupTriggerPropertiesRequestInfo1 struct{}"
	}

	return strings.Join([]string{"BackupTriggerPropertiesRequestInfo1", string(data)}, " ")
}
