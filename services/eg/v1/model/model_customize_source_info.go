package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CustomizeSourceInfo struct {

	// 事件源ID
	Id *string `json:"id,omitempty"`

	// 事件源名称
	Name *string `json:"name,omitempty"`

	// 事件源名称展示
	Label *string `json:"label,omitempty"`

	// 事件源描述
	Description *string `json:"description,omitempty"`

	// 事件源提供方类型，OFFICIAL：官方云服务事件源；CUSTOM：用户创建的自定义事件源；PARTNER：伙伴事件源
	ProviderType *CustomizeSourceInfoProviderType `json:"provider_type,omitempty"`

	// 事件源提供的事件类型列表，只有官方云服务事件源提供事件类型
	EventTypes *[]CustomizeSourceInfoEventTypes `json:"event_types,omitempty"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新UTC时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 事件源归属的事件通道ID
	ChannelId *string `json:"channel_id,omitempty"`

	// 事件源归属的事件通道名称
	ChannelName *string `json:"channel_name,omitempty"`

	// 事件源类型
	Type *string `json:"type,omitempty"`

	// json格式封装消息实例链接信息：如RabbitMQ实例的instance_id字段、虚拟主机vhost字段、队列queue字段、用户名、密码等
	Detail *interface{} `json:"detail,omitempty"`

	// 自定义事件源状态
	Status *CustomizeSourceInfoStatus `json:"status,omitempty"`

	ErrorInfo *ErrorInfo `json:"error_info,omitempty"`
}

func (o CustomizeSourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeSourceInfo struct{}"
	}

	return strings.Join([]string{"CustomizeSourceInfo", string(data)}, " ")
}

type CustomizeSourceInfoProviderType struct {
	value string
}

type CustomizeSourceInfoProviderTypeEnum struct {
	OFFICIAL CustomizeSourceInfoProviderType
	CUSTOM   CustomizeSourceInfoProviderType
	PARTNER  CustomizeSourceInfoProviderType
}

func GetCustomizeSourceInfoProviderTypeEnum() CustomizeSourceInfoProviderTypeEnum {
	return CustomizeSourceInfoProviderTypeEnum{
		OFFICIAL: CustomizeSourceInfoProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: CustomizeSourceInfoProviderType{
			value: "CUSTOM",
		},
		PARTNER: CustomizeSourceInfoProviderType{
			value: "PARTNER",
		},
	}
}

func (c CustomizeSourceInfoProviderType) Value() string {
	return c.value
}

func (c CustomizeSourceInfoProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomizeSourceInfoProviderType) UnmarshalJSON(b []byte) error {
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

type CustomizeSourceInfoStatus struct {
	value string
}

type CustomizeSourceInfoStatusEnum struct {
	CREATE_FAILED CustomizeSourceInfoStatus
	RUNNING       CustomizeSourceInfoStatus
	ERROR         CustomizeSourceInfoStatus
}

func GetCustomizeSourceInfoStatusEnum() CustomizeSourceInfoStatusEnum {
	return CustomizeSourceInfoStatusEnum{
		CREATE_FAILED: CustomizeSourceInfoStatus{
			value: "CREATE_FAILED",
		},
		RUNNING: CustomizeSourceInfoStatus{
			value: "RUNNING",
		},
		ERROR: CustomizeSourceInfoStatus{
			value: "ERROR",
		},
	}
}

func (c CustomizeSourceInfoStatus) Value() string {
	return c.value
}

func (c CustomizeSourceInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomizeSourceInfoStatus) UnmarshalJSON(b []byte) error {
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
