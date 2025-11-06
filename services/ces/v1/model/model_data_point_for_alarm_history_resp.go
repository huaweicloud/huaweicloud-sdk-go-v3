package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataPointForAlarmHistoryResp **参数解释** 计算出该条告警历史的资源监控数据上报时间和监控数值。
type DataPointForAlarmHistoryResp struct {

	// **参数解释** 计算出该条告警历史的资源监控数据上报时间，UNIX时间戳，单位毫秒，如：1603131028000。 **取值范围**： 不涉及
	Time *int64 `json:"time,omitempty"`

	// **参数解释** 计算出该条告警历史的资源监控数据在该时间点的监控数值，如：7.019。 **取值范围**： 具体阈值取值请参见附录中各服务监控指标中取值范围，如“[支持监控的服务列表](ces_03_0059.xml)”中ECS的CPU使用率cpu_util取值范围可配置80。最小值为0，最大值为1.7976931348623157e+108。
	Value *float64 `json:"value,omitempty"`
}

func (o DataPointForAlarmHistoryResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataPointForAlarmHistoryResp struct{}"
	}

	return strings.Join([]string{"DataPointForAlarmHistoryResp", string(data)}, " ")
}
