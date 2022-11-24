package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateEventSourceResponse struct {

	// 事件源ID
	Id *string `json:"id,omitempty"`

	// 事件源名称
	Name *string `json:"name,omitempty"`

	// 事件源名称展示
	Label *string `json:"label,omitempty"`

	// 事件源描述
	Description *string `json:"description,omitempty"`

	// 事件源提供方类型，OFFICIAL：官方云服务事件源；CUSTOM：用户创建的自定义事件源
	ProviderType *UpdateEventSourceResponseProviderType `json:"provider_type,omitempty"`

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
	Status *UpdateEventSourceResponseStatus `json:"status,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEventSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEventSourceResponse struct{}"
	}

	return strings.Join([]string{"UpdateEventSourceResponse", string(data)}, " ")
}

type UpdateEventSourceResponseProviderType struct {
	value string
}

type UpdateEventSourceResponseProviderTypeEnum struct {
	OFFICIAL UpdateEventSourceResponseProviderType
	CUSTOM   UpdateEventSourceResponseProviderType
}

func GetUpdateEventSourceResponseProviderTypeEnum() UpdateEventSourceResponseProviderTypeEnum {
	return UpdateEventSourceResponseProviderTypeEnum{
		OFFICIAL: UpdateEventSourceResponseProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: UpdateEventSourceResponseProviderType{
			value: "CUSTOM",
		},
	}
}

func (c UpdateEventSourceResponseProviderType) Value() string {
	return c.value
}

func (c UpdateEventSourceResponseProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEventSourceResponseProviderType) UnmarshalJSON(b []byte) error {
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

type UpdateEventSourceResponseStatus struct {
	value string
}

type UpdateEventSourceResponseStatusEnum struct {
	CREATE_FAILED UpdateEventSourceResponseStatus
	RUNNING       UpdateEventSourceResponseStatus
	ERROR         UpdateEventSourceResponseStatus
}

func GetUpdateEventSourceResponseStatusEnum() UpdateEventSourceResponseStatusEnum {
	return UpdateEventSourceResponseStatusEnum{
		CREATE_FAILED: UpdateEventSourceResponseStatus{
			value: "CREATE_FAILED",
		},
		RUNNING: UpdateEventSourceResponseStatus{
			value: "RUNNING",
		},
		ERROR: UpdateEventSourceResponseStatus{
			value: "ERROR",
		},
	}
}

func (c UpdateEventSourceResponseStatus) Value() string {
	return c.value
}

func (c UpdateEventSourceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEventSourceResponseStatus) UnmarshalJSON(b []byte) error {
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
