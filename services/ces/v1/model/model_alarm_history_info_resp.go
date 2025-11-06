package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmHistoryInfoResp **参数解释**： 告警历史详细信息
type AlarmHistoryInfoResp struct {

	// **参数解释**： 告警规则的ID，如：al1603131199286dzxpqK3Ez。 **取值范围**： 字符串长度为24
	AlarmId *string `json:"alarm_id,omitempty"`

	// **参数解释**： 告警规则的名称，如：alarm-test01 **取值范围**： 字符串长度在 1 到 128 之间
	AlarmName *string `json:"alarm_name,omitempty"`

	// **参数解释**： 告警规则的描述 **取值范围**： 字符串长度在 0 到 256 之间
	AlarmDescription *string `json:"alarm_description,omitempty"`

	Metric *MetricInfoResp `json:"metric,omitempty"`

	Condition *ConditionResp `json:"condition,omitempty"`

	// **参数解释**： 告警记录的告警级别。 **取值范围**： 枚举值： - 1：紧急 - 2：重要 - 3：次要 - 4：提示
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// **参数解释**： 告警规则类型 **取值范围**： 枚举值: - ALL_INSTANCE：全部资源指标告警 - RESOURCE_GROUP：资源分组指标告警 - MULTI_INSTANCE：指定资源指标告警 - EVENT.SYS：系统事件告警 - EVENT.CUSTOM：自定义事件告警 - DNSHealthCheck：健康检查告警
	AlarmType *string `json:"alarm_type,omitempty"`

	// **参数解释**： 告警规则是否被启用 **取值范围**： 值为true或者false - true：开启 - false：关闭
	AlarmEnabled *bool `json:"alarm_enabled,omitempty"`

	// **参数解释**： 是否发送通知 **取值范围**： 值为true或者false - true：发送通知 - false：不发送通知
	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`

	// **参数解释**： 告警触发的动作。结构如下：{  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  }  type取值：notification：通知。autoscaling：弹性伸缩。notificationList：告警状态发生变化时，被通知对象的列表。
	AlarmActions *[][]NotificationResp `json:"alarm_actions,omitempty"`

	// **参数解释**： 告警恢复触发的动作。结构如下：{  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  } type取值：notification：通知。notificationList：告警状态发生变化时，被通知对象的列表。
	OkActions *[][]NotificationResp `json:"ok_actions,omitempty"`

	// **参数解释**： 数据不足触发的动作。结构如下：{  \"type\": \"notification\", \"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"]  } type取值：数据不足触发告警通知类型，取值为notification。notificationList：数据不足触发告警通知时，被通知对象的ID列表。
	InsufficientdataActions *[][]NotificationResp `json:"insufficientdata_actions,omitempty"`

	// **参数解释**： 告警状态变更的时间，UNIX时间戳，单位毫秒，如：1603131199000 **取值范围**： 不涉及
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**： 企业项目ID **取值范围**： 只能包含小写字母、数字、“-”、“_”，可以自定义企业项目ID，长度为36个字符。也可以为0（代表默认企业项目ID），all_granted_eps（代表所有企业项目ID）
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 计算出该条告警历史的时间，UNIX时间戳，单位毫秒，如：1603131199469 **取值范围**： 不涉及
	TriggerTime *int64 `json:"trigger_time,omitempty"`

	// **参数解释**： 告警历史的状态 **取值范围**： 枚举值： - ok：正常 - alarm：告警 - insufficient_data：数据不足 - invalid：已失效
	AlarmStatus *string `json:"alarm_status,omitempty"`

	// **参数解释**： 计算出该条告警历史的资源监控数据的一组数据上报时间和监控数值
	Datapoints *[]DataPointForAlarmHistoryResp `json:"datapoints,omitempty"`

	AdditionalInfo *AdditionalInfoResp `json:"additional_info,omitempty"`

	// **参数解释** 通知方式 **取值范围**： 枚举值： - NOTIFICATION_POLICY：通知策略 - NOTIFICATION_GROUP：通知组 - TOPIC_SUBSCRIPTION：主题订阅
	NotificationManner *AlarmHistoryInfoRespNotificationManner `json:"notification_manner,omitempty"`
}

func (o AlarmHistoryInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryInfoResp struct{}"
	}

	return strings.Join([]string{"AlarmHistoryInfoResp", string(data)}, " ")
}

type AlarmHistoryInfoRespNotificationManner struct {
	value string
}

type AlarmHistoryInfoRespNotificationMannerEnum struct {
	NOTIFICATION_POLICY AlarmHistoryInfoRespNotificationManner
	NOTIFICATION_GROUP  AlarmHistoryInfoRespNotificationManner
	TOPIC_SUBSCRIPTION  AlarmHistoryInfoRespNotificationManner
}

func GetAlarmHistoryInfoRespNotificationMannerEnum() AlarmHistoryInfoRespNotificationMannerEnum {
	return AlarmHistoryInfoRespNotificationMannerEnum{
		NOTIFICATION_POLICY: AlarmHistoryInfoRespNotificationManner{
			value: "NOTIFICATION_POLICY",
		},
		NOTIFICATION_GROUP: AlarmHistoryInfoRespNotificationManner{
			value: "NOTIFICATION_GROUP",
		},
		TOPIC_SUBSCRIPTION: AlarmHistoryInfoRespNotificationManner{
			value: "TOPIC_SUBSCRIPTION",
		},
	}
}

func (c AlarmHistoryInfoRespNotificationManner) Value() string {
	return c.value
}

func (c AlarmHistoryInfoRespNotificationManner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmHistoryInfoRespNotificationManner) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
