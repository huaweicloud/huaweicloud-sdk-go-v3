package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListAlarmResponseAlarms struct {

	// 告警规则id，以al开头，包含22个数字或字母
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128
	Name *string `json:"name,omitempty"`

	// 告警描述，长度0-256
	Description *string `json:"description,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考“[服务命名名称](ces_03_0059.xml)”
	Namespace *string `json:"namespace,omitempty"`

	// 告警策略
	Policies *[]Policy `json:"policies,omitempty"`

	// 资源列表，关联资源需要使用查询告警规则资源接口获取
	Resources *[]ResourcesInListResp `json:"resources,omitempty"`

	Type *AlarmType `json:"type,omitempty"`

	// 是否开启告警规则。true:开启，false:关闭。
	Enabled *bool `json:"enabled,omitempty"`

	// 是否开启告警通知。true:开启，false:关闭。
	NotificationEnabled *bool `json:"notification_enabled,omitempty"`

	// 告警触发的动作
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// 告警恢复触发的动作
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// 告警通知开启时间
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// 告警通知关闭时间
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"
	EffectiveTimezone *string `json:"effective_timezone,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 告警规则关联告警模板ID，如果传了，告警规则关联的策略会和告警模板策略联动变化
	AlarmTemplateId *string `json:"alarm_template_id,omitempty"`

	// 产品层级跨纬规则需要指明的规则产品名称，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"
	ProductName *string `json:"product_name,omitempty"`

	// 产品层级跨纬规则需要指明为产品层级规则，resource_level取值为product即为产品层级跨纬规则，不填或者取值为dimension则为旧的规则类型
	ResourceLevel *ListAlarmResponseAlarmsResourceLevel `json:"resource_level,omitempty"`
}

func (o ListAlarmResponseAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmResponseAlarms struct{}"
	}

	return strings.Join([]string{"ListAlarmResponseAlarms", string(data)}, " ")
}

type ListAlarmResponseAlarmsResourceLevel struct {
	value string
}

type ListAlarmResponseAlarmsResourceLevelEnum struct {
	PRODUCT   ListAlarmResponseAlarmsResourceLevel
	DIMENSION ListAlarmResponseAlarmsResourceLevel
}

func GetListAlarmResponseAlarmsResourceLevelEnum() ListAlarmResponseAlarmsResourceLevelEnum {
	return ListAlarmResponseAlarmsResourceLevelEnum{
		PRODUCT: ListAlarmResponseAlarmsResourceLevel{
			value: "product",
		},
		DIMENSION: ListAlarmResponseAlarmsResourceLevel{
			value: "dimension",
		},
	}
}

func (c ListAlarmResponseAlarmsResourceLevel) Value() string {
	return c.value
}

func (c ListAlarmResponseAlarmsResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmResponseAlarmsResourceLevel) UnmarshalJSON(b []byte) error {
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
