package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateSubscriptionResponse struct {

	// 事件订阅ID
	Id *string `json:"id,omitempty"`

	// 事件订阅名称
	Name *string `json:"name,omitempty"`

	// 事件订阅描述
	Description *string `json:"description,omitempty"`

	// 事件订阅类型
	Type *UpdateSubscriptionResponseType `json:"type,omitempty"`

	// 事件订阅状态
	Status *UpdateSubscriptionResponseStatus `json:"status,omitempty"`

	// 通道ID
	ChannelId *string `json:"channel_id,omitempty"`

	// 通道名称
	ChannelName *string `json:"channel_name,omitempty"`

	// 订阅源列表
	Sources *[]SubscriptionSourceInfo `json:"sources,omitempty"`

	// 订阅目标列表
	Targets *[]SubscriptionTargetInfo `json:"targets,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime    *string `json:"updated_time,omitempty"`
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
