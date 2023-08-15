package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSubscriptionResponse Response Object
type UpdateSubscriptionResponse struct {

	// 订阅ID
	Id *string `json:"id,omitempty"`

	// 订阅名称
	Name *string `json:"name,omitempty"`

	// 订阅描述
	Description *string `json:"description,omitempty"`

	// 类型
	Type *UpdateSubscriptionResponseType `json:"type,omitempty"`

	// 状态
	Status *UpdateSubscriptionResponseStatus `json:"status,omitempty"`

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

func (o UpdateSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionResponse", string(data)}, " ")
}

type UpdateSubscriptionResponseType struct {
	value string
}

type UpdateSubscriptionResponseTypeEnum struct {
	EVENT     UpdateSubscriptionResponseType
	SCHEDULED UpdateSubscriptionResponseType
}

func GetUpdateSubscriptionResponseTypeEnum() UpdateSubscriptionResponseTypeEnum {
	return UpdateSubscriptionResponseTypeEnum{
		EVENT: UpdateSubscriptionResponseType{
			value: "EVENT",
		},
		SCHEDULED: UpdateSubscriptionResponseType{
			value: "SCHEDULED",
		},
	}
}

func (c UpdateSubscriptionResponseType) Value() string {
	return c.value
}

func (c UpdateSubscriptionResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSubscriptionResponseType) UnmarshalJSON(b []byte) error {
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

type UpdateSubscriptionResponseStatus struct {
	value string
}

type UpdateSubscriptionResponseStatusEnum struct {
	CREATED  UpdateSubscriptionResponseStatus
	ENABLED  UpdateSubscriptionResponseStatus
	DISABLED UpdateSubscriptionResponseStatus
	FROZEN   UpdateSubscriptionResponseStatus
	ERROR    UpdateSubscriptionResponseStatus
}

func GetUpdateSubscriptionResponseStatusEnum() UpdateSubscriptionResponseStatusEnum {
	return UpdateSubscriptionResponseStatusEnum{
		CREATED: UpdateSubscriptionResponseStatus{
			value: "CREATED",
		},
		ENABLED: UpdateSubscriptionResponseStatus{
			value: "ENABLED",
		},
		DISABLED: UpdateSubscriptionResponseStatus{
			value: "DISABLED",
		},
		FROZEN: UpdateSubscriptionResponseStatus{
			value: "FROZEN",
		},
		ERROR: UpdateSubscriptionResponseStatus{
			value: "ERROR",
		},
	}
}

func (c UpdateSubscriptionResponseStatus) Value() string {
	return c.value
}

func (c UpdateSubscriptionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSubscriptionResponseStatus) UnmarshalJSON(b []byte) error {
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
