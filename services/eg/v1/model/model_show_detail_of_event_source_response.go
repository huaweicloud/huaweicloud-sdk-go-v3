package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowDetailOfEventSourceResponse struct {

	// 事件源ID
	Id *string `json:"id,omitempty"`

	// 事件源名称
	Name *string `json:"name,omitempty"`

	// 事件源名称展示
	Label *string `json:"label,omitempty"`

	// 事件源描述
	Description *string `json:"description,omitempty"`

	// 事件源提供方类型，OFFICIAL：官方云服务事件源；CUSTOM：用户创建的自定义事件源；PARTNER：伙伴事件源
	ProviderType *ShowDetailOfEventSourceResponseProviderType `json:"provider_type,omitempty"`

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
	Status         *ShowDetailOfEventSourceResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ShowDetailOfEventSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfEventSourceResponse struct{}"
	}

	return strings.Join([]string{"ShowDetailOfEventSourceResponse", string(data)}, " ")
}

type ShowDetailOfEventSourceResponseProviderType struct {
	value string
}

type ShowDetailOfEventSourceResponseProviderTypeEnum struct {
	OFFICIAL ShowDetailOfEventSourceResponseProviderType
	CUSTOM   ShowDetailOfEventSourceResponseProviderType
	PARTNER  ShowDetailOfEventSourceResponseProviderType
}

func GetShowDetailOfEventSourceResponseProviderTypeEnum() ShowDetailOfEventSourceResponseProviderTypeEnum {
	return ShowDetailOfEventSourceResponseProviderTypeEnum{
		OFFICIAL: ShowDetailOfEventSourceResponseProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: ShowDetailOfEventSourceResponseProviderType{
			value: "CUSTOM",
		},
		PARTNER: ShowDetailOfEventSourceResponseProviderType{
			value: "PARTNER",
		},
	}
}

func (c ShowDetailOfEventSourceResponseProviderType) Value() string {
	return c.value
}

func (c ShowDetailOfEventSourceResponseProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailOfEventSourceResponseProviderType) UnmarshalJSON(b []byte) error {
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

type ShowDetailOfEventSourceResponseStatus struct {
	value string
}

type ShowDetailOfEventSourceResponseStatusEnum struct {
	CREATE_FAILED ShowDetailOfEventSourceResponseStatus
	RUNNING       ShowDetailOfEventSourceResponseStatus
	ERROR         ShowDetailOfEventSourceResponseStatus
}

func GetShowDetailOfEventSourceResponseStatusEnum() ShowDetailOfEventSourceResponseStatusEnum {
	return ShowDetailOfEventSourceResponseStatusEnum{
		CREATE_FAILED: ShowDetailOfEventSourceResponseStatus{
			value: "CREATE_FAILED",
		},
		RUNNING: ShowDetailOfEventSourceResponseStatus{
			value: "RUNNING",
		},
		ERROR: ShowDetailOfEventSourceResponseStatus{
			value: "ERROR",
		},
	}
}

func (c ShowDetailOfEventSourceResponseStatus) Value() string {
	return c.value
}

func (c ShowDetailOfEventSourceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailOfEventSourceResponseStatus) UnmarshalJSON(b []byte) error {
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
