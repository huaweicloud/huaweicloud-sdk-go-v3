package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PostAlarmsReqV2 struct {

	// 告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128
	Name string `json:"name"`

	// 告警描述，长度0-256
	Description *string `json:"description,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考“[服务维度名称](ces_03_0059.xml)”
	Namespace string `json:"namespace"`

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	ResourceGroupId *string `json:"resource_group_id,omitempty"`

	// 资源列表，告警规则类型为全部资源、资源分组时，资源维度值传空；告警规则类型为指定资源时，资源维度值必填，可以同时指定监控多个资源。
	Resources [][]Dimension `json:"resources"`

	// 告警策略，当alarm_template_id字段为空时必填，不为空时不填
	Policies *[]Policy `json:"policies,omitempty"`

	Type *AlarmType `json:"type"`

	// 告警触发的动作
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// 告警恢复触发的动作
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// 告警通知开启时间
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// 告警通知关闭时间
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否开启告警规则。true:开启，false:关闭。
	Enabled bool `json:"enabled"`

	// 是否开启告警通知。true:开启，false:关闭。
	NotificationEnabled bool `json:"notification_enabled"`

	// 告警规则关联告警模板ID，如果传了，告警规则关联的策略会和告警模板策略联动变化
	AlarmTemplateId *string `json:"alarm_template_id,omitempty"`

	// 租户标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 产品层级跨纬规则创建时需要指明的规则产品名称，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"
	ProductName *string `json:"product_name,omitempty"`

	// 产品层级跨纬规则创建时需要指明为产品层级规则，resource_level取值为product即为产品层级跨纬规则，不填或者取值为dimension则为旧的规则类型
	ResourceLevel *PostAlarmsReqV2ResourceLevel `json:"resource_level,omitempty"`
}

func (o PostAlarmsReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostAlarmsReqV2 struct{}"
	}

	return strings.Join([]string{"PostAlarmsReqV2", string(data)}, " ")
}

type PostAlarmsReqV2ResourceLevel struct {
	value string
}

type PostAlarmsReqV2ResourceLevelEnum struct {
	PRODUCT   PostAlarmsReqV2ResourceLevel
	DIMENSION PostAlarmsReqV2ResourceLevel
}

func GetPostAlarmsReqV2ResourceLevelEnum() PostAlarmsReqV2ResourceLevelEnum {
	return PostAlarmsReqV2ResourceLevelEnum{
		PRODUCT: PostAlarmsReqV2ResourceLevel{
			value: "product",
		},
		DIMENSION: PostAlarmsReqV2ResourceLevel{
			value: "dimension",
		},
	}
}

func (c PostAlarmsReqV2ResourceLevel) Value() string {
	return c.value
}

func (c PostAlarmsReqV2ResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostAlarmsReqV2ResourceLevel) UnmarshalJSON(b []byte) error {
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
