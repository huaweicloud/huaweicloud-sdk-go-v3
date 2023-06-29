package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDetailOfSubscriptionResponse Response Object
type ShowDetailOfSubscriptionResponse struct {

	// 订阅ID
	Id *string `json:"id,omitempty"`

	// 订阅名称
	Name *string `json:"name,omitempty"`

	// 订阅描述
	Description *string `json:"description,omitempty"`

	// 类型
	Type *ShowDetailOfSubscriptionResponseType `json:"type,omitempty"`

	// 状态
	Status *ShowDetailOfSubscriptionResponseStatus `json:"status,omitempty"`

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
	UpdatedTime    *string `json:"updated_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDetailOfSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"ShowDetailOfSubscriptionResponse", string(data)}, " ")
}

type ShowDetailOfSubscriptionResponseType struct {
	value string
}

type ShowDetailOfSubscriptionResponseTypeEnum struct {
	EVENT     ShowDetailOfSubscriptionResponseType
	SCHEDULED ShowDetailOfSubscriptionResponseType
}

func GetShowDetailOfSubscriptionResponseTypeEnum() ShowDetailOfSubscriptionResponseTypeEnum {
	return ShowDetailOfSubscriptionResponseTypeEnum{
		EVENT: ShowDetailOfSubscriptionResponseType{
			value: "EVENT",
		},
		SCHEDULED: ShowDetailOfSubscriptionResponseType{
			value: "SCHEDULED",
		},
	}
}

func (c ShowDetailOfSubscriptionResponseType) Value() string {
	return c.value
}

func (c ShowDetailOfSubscriptionResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailOfSubscriptionResponseType) UnmarshalJSON(b []byte) error {
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

type ShowDetailOfSubscriptionResponseStatus struct {
	value string
}

type ShowDetailOfSubscriptionResponseStatusEnum struct {
	CREATED  ShowDetailOfSubscriptionResponseStatus
	ENABLED  ShowDetailOfSubscriptionResponseStatus
	DISABLED ShowDetailOfSubscriptionResponseStatus
	FROZEN   ShowDetailOfSubscriptionResponseStatus
	ERROR    ShowDetailOfSubscriptionResponseStatus
}

func GetShowDetailOfSubscriptionResponseStatusEnum() ShowDetailOfSubscriptionResponseStatusEnum {
	return ShowDetailOfSubscriptionResponseStatusEnum{
		CREATED: ShowDetailOfSubscriptionResponseStatus{
			value: "CREATED",
		},
		ENABLED: ShowDetailOfSubscriptionResponseStatus{
			value: "ENABLED",
		},
		DISABLED: ShowDetailOfSubscriptionResponseStatus{
			value: "DISABLED",
		},
		FROZEN: ShowDetailOfSubscriptionResponseStatus{
			value: "FROZEN",
		},
		ERROR: ShowDetailOfSubscriptionResponseStatus{
			value: "ERROR",
		},
	}
}

func (c ShowDetailOfSubscriptionResponseStatus) Value() string {
	return c.value
}

func (c ShowDetailOfSubscriptionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailOfSubscriptionResponseStatus) UnmarshalJSON(b []byte) error {
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
