package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SubscriptionInfo struct {

	// 订阅ID
	Id *string `json:"id,omitempty"`

	// 订阅名称
	Name *string `json:"name,omitempty"`

	// 订阅描述
	Description *string `json:"description,omitempty"`

	// 类型
	Type *SubscriptionInfoType `json:"type,omitempty"`

	// 状态
	Status *SubscriptionInfoStatus `json:"status,omitempty"`

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
}

func (o SubscriptionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionInfo struct{}"
	}

	return strings.Join([]string{"SubscriptionInfo", string(data)}, " ")
}

type SubscriptionInfoType struct {
	value string
}

type SubscriptionInfoTypeEnum struct {
	EVENT     SubscriptionInfoType
	SCHEDULED SubscriptionInfoType
}

func GetSubscriptionInfoTypeEnum() SubscriptionInfoTypeEnum {
	return SubscriptionInfoTypeEnum{
		EVENT: SubscriptionInfoType{
			value: "EVENT",
		},
		SCHEDULED: SubscriptionInfoType{
			value: "SCHEDULED",
		},
	}
}

func (c SubscriptionInfoType) Value() string {
	return c.value
}

func (c SubscriptionInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubscriptionInfoType) UnmarshalJSON(b []byte) error {
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

type SubscriptionInfoStatus struct {
	value string
}

type SubscriptionInfoStatusEnum struct {
	CREATED  SubscriptionInfoStatus
	ENABLED  SubscriptionInfoStatus
	DISABLED SubscriptionInfoStatus
	FROZEN   SubscriptionInfoStatus
	ERROR    SubscriptionInfoStatus
}

func GetSubscriptionInfoStatusEnum() SubscriptionInfoStatusEnum {
	return SubscriptionInfoStatusEnum{
		CREATED: SubscriptionInfoStatus{
			value: "CREATED",
		},
		ENABLED: SubscriptionInfoStatus{
			value: "ENABLED",
		},
		DISABLED: SubscriptionInfoStatus{
			value: "DISABLED",
		},
		FROZEN: SubscriptionInfoStatus{
			value: "FROZEN",
		},
		ERROR: SubscriptionInfoStatus{
			value: "ERROR",
		},
	}
}

func (c SubscriptionInfoStatus) Value() string {
	return c.value
}

func (c SubscriptionInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubscriptionInfoStatus) UnmarshalJSON(b []byte) error {
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
