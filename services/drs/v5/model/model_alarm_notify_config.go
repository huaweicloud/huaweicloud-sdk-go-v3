package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmNotifyConfig 异常通知信息设置信息体。配置该参数时，代表当任务状态异常时，系统将发送通知给指定的SMN Topic。
type AlarmNotifyConfig struct {

	// 异常告警是否通知用户。
	AlarmToUser bool `json:"alarm_to_user"`

	// SMN主题URN。
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 时延阈值(单位为s)。取值： - 最小值：1 - 最大值：3600 - 缺省值：0 - 说明： 1.源数据库和目标数据库之间的同步有时会存在一个时间差，称为时延，单位为秒。 2.时延阈值设置是指时延超过一定的值并持续6分钟后，DRS可以发送通知给指定收件人。（初次进入增量迁移阶段，会有较多数据等待同步，存在较大的时延，属于正常情况，不在此功能的监控范围之内。） 3.只有全量+增量的任务支持此选项。
	DelayTime *int64 `json:"delay_time,omitempty"`

	// RPO时延阈值(单位为s)。取值： - 最小值：1 - 最大值：3600 - 缺省值：0 - 说明： RPO时延阈值设置是业务数据库与DRS实例间同步的时延超过一定的值并持续6分钟后，DRS可以发送通知给指定收件人。（初次进入增量灾备阶段，会有较多数据等待同步，存在较大的时延，属于正常情况，不在此功能的监控范围之内。）
	RpoDelay *int64 `json:"rpo_delay,omitempty"`

	// RTO时延阈值(s)。取值： - 最小值：1 - 最大值：3600 - 缺省值：0 - 说明： RTO时延阈值设置是DRS实例与灾备数据库间同步的时延超过一定的值并持续6分钟后，DRS可以发送通知给指定收件人。
	RtoDelay *int64 `json:"rto_delay,omitempty"`
}

func (o AlarmNotifyConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmNotifyConfig struct{}"
	}

	return strings.Join([]string{"AlarmNotifyConfig", string(data)}, " ")
}
