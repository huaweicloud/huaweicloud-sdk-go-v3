package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSubscriptionResponse Response Object
type CreateSubscriptionResponse struct {

	// 订阅ID
	Id *string `json:"id,omitempty"`

	// 订阅名称
	Name *string `json:"name,omitempty"`

	// 订阅描述
	Description *string `json:"description,omitempty"`

	// 类型
	Type *CreateSubscriptionResponseType `json:"type,omitempty"`

	// 状态
	Status *CreateSubscriptionResponseStatus `json:"status,omitempty"`

	// 通道ID
	ChannelId *string `json:"channel_id,omitempty"`

	// 通道名称
	ChannelName *string `json:"channel_name,omitempty"`

	// 标签信息
	Used *[]SubscriptionUsedInfo `json:"used,omitempty"`

	// 订阅源列表
	Sources *[]SubscriptionSourceInfo `json:"sources,omitempty"`

	// 订阅目标列表
	Targets *[]SubscriptionTargetInfo `json:"targets,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionResponse", string(data)}, " ")
}

type CreateSubscriptionResponseType struct {
	value string
}

type CreateSubscriptionResponseTypeEnum struct {
	EVENT     CreateSubscriptionResponseType
	SCHEDULED CreateSubscriptionResponseType
}

func GetCreateSubscriptionResponseTypeEnum() CreateSubscriptionResponseTypeEnum {
	return CreateSubscriptionResponseTypeEnum{
		EVENT: CreateSubscriptionResponseType{
			value: "EVENT",
		},
		SCHEDULED: CreateSubscriptionResponseType{
			value: "SCHEDULED",
		},
	}
}

func (c CreateSubscriptionResponseType) Value() string {
	return c.value
}

func (c CreateSubscriptionResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSubscriptionResponseType) UnmarshalJSON(b []byte) error {
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

type CreateSubscriptionResponseStatus struct {
	value string
}

type CreateSubscriptionResponseStatusEnum struct {
	CREATED  CreateSubscriptionResponseStatus
	ENABLED  CreateSubscriptionResponseStatus
	DISABLED CreateSubscriptionResponseStatus
	FROZEN   CreateSubscriptionResponseStatus
	ERROR    CreateSubscriptionResponseStatus
}

func GetCreateSubscriptionResponseStatusEnum() CreateSubscriptionResponseStatusEnum {
	return CreateSubscriptionResponseStatusEnum{
		CREATED: CreateSubscriptionResponseStatus{
			value: "CREATED",
		},
		ENABLED: CreateSubscriptionResponseStatus{
			value: "ENABLED",
		},
		DISABLED: CreateSubscriptionResponseStatus{
			value: "DISABLED",
		},
		FROZEN: CreateSubscriptionResponseStatus{
			value: "FROZEN",
		},
		ERROR: CreateSubscriptionResponseStatus{
			value: "ERROR",
		},
	}
}

func (c CreateSubscriptionResponseStatus) Value() string {
	return c.value
}

func (c CreateSubscriptionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSubscriptionResponseStatus) UnmarshalJSON(b []byte) error {
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
