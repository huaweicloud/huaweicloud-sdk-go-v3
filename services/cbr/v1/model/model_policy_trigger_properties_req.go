package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyTriggerPropertiesReq
type PolicyTriggerPropertiesReq struct {

	// 调度规则。限制24条规则。调度器的调度规则，可参照iCalendar RFC 2445规范中的事件规则，但仅支持FREQ、BYDAY、BYHOUR、BYMINUTE、INTERVAL等参数，其中FREQ仅支持WEEKLY和DAILY，BYDAY支持一周七天（MO、TU、WE、TH、FR、SA、SU），BYHOUR支持0-23小时，BYMINUTE支持0-59分钟，并且间隔不能小于一小时，一天最大24个时间点。例如，周一到周天，每天UTC时间的14:00调度，其规则为：'FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR,SA,SU;BYHOUR=14;BYMINUTE=00'。例如：某个地区的时间为 UTC+8，如果在该地区每天14:00调度，则在14点基础上减8，其规则为'FREQ=DAILY;INTERVAL=1;BYHOUR=6;BYMINUTE=00'。
	Pattern []string `json:"pattern"`
}

func (o PolicyTriggerPropertiesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTriggerPropertiesReq struct{}"
	}

	return strings.Join([]string{"PolicyTriggerPropertiesReq", string(data)}, " ")
}
