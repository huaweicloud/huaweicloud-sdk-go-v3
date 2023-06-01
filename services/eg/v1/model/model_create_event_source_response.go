package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateEventSourceResponse struct {

	// 事件源ID
	Id *string `json:"id,omitempty"`

	// 事件源名称
	Name *string `json:"name,omitempty"`

	// 事件源名称展示
	Label *string `json:"label,omitempty"`

	// 事件源描述
	Description *string `json:"description,omitempty"`

	// 事件源提供方类型，OFFICIAL：官方云服务事件源；CUSTOM：用户创建的自定义事件源；PARTNER：伙伴事件源
	ProviderType *CreateEventSourceResponseProviderType `json:"provider_type,omitempty"`

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
	Status *CreateEventSourceResponseStatus `json:"status,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEventSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventSourceResponse struct{}"
	}

	return strings.Join([]string{"CreateEventSourceResponse", string(data)}, " ")
}

type CreateEventSourceResponseProviderType struct {
	value string
}

type CreateEventSourceResponseProviderTypeEnum struct {
	OFFICIAL CreateEventSourceResponseProviderType
	CUSTOM   CreateEventSourceResponseProviderType
	PARTNER  CreateEventSourceResponseProviderType
}

func GetCreateEventSourceResponseProviderTypeEnum() CreateEventSourceResponseProviderTypeEnum {
	return CreateEventSourceResponseProviderTypeEnum{
		OFFICIAL: CreateEventSourceResponseProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: CreateEventSourceResponseProviderType{
			value: "CUSTOM",
		},
		PARTNER: CreateEventSourceResponseProviderType{
			value: "PARTNER",
		},
	}
}

func (c CreateEventSourceResponseProviderType) Value() string {
	return c.value
}

func (c CreateEventSourceResponseProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEventSourceResponseProviderType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateEventSourceResponseStatus struct {
	value string
}

type CreateEventSourceResponseStatusEnum struct {
	CREATE_FAILED CreateEventSourceResponseStatus
	RUNNING       CreateEventSourceResponseStatus
	ERROR         CreateEventSourceResponseStatus
}

func GetCreateEventSourceResponseStatusEnum() CreateEventSourceResponseStatusEnum {
	return CreateEventSourceResponseStatusEnum{
		CREATE_FAILED: CreateEventSourceResponseStatus{
			value: "CREATE_FAILED",
		},
		RUNNING: CreateEventSourceResponseStatus{
			value: "RUNNING",
		},
		ERROR: CreateEventSourceResponseStatus{
			value: "ERROR",
		},
	}
}

func (c CreateEventSourceResponseStatus) Value() string {
	return c.value
}

func (c CreateEventSourceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEventSourceResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
