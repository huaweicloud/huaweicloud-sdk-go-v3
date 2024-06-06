package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateNoticeRuleItem struct {

	// 通知名称。
	Name string `json:"name"`

	// 触发事件名称，支持实例调度成功、实例调度失败、健康检查成功、健康检查失败、镜像拉取成功、镜像拉取失败、容器启动成功、容器启动失败、卷挂载成功、卷挂载失败。
	EventName *CreateNoticeRuleItemEventName `json:"event_name,omitempty"`

	Scope *NoticeRuleScope `json:"scope"`

	TriggerPolicy *TriggerPolicy `json:"trigger_policy"`

	Notification *NoticeRuleNotification `json:"notification"`

	// 是否启用。
	Enable *bool `json:"enable,omitempty"`
}

func (o CreateNoticeRuleItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNoticeRuleItem struct{}"
	}

	return strings.Join([]string{"CreateNoticeRuleItem", string(data)}, " ")
}

type CreateNoticeRuleItemEventName struct {
	value string
}

type CreateNoticeRuleItemEventNameEnum struct {
	SCHEDULED               CreateNoticeRuleItemEventName
	FAILED_SCHEDULING       CreateNoticeRuleItemEventName
	HEALTHY                 CreateNoticeRuleItemEventName
	UNHEALTHY               CreateNoticeRuleItemEventName
	PULLED                  CreateNoticeRuleItemEventName
	FAILED_PULL_IMAGE       CreateNoticeRuleItemEventName
	STARTED                 CreateNoticeRuleItemEventName
	BACK_OFF_START          CreateNoticeRuleItemEventName
	SUCCESSFUL_MOUNT_VOLUME CreateNoticeRuleItemEventName
	FAILED_MOUNT            CreateNoticeRuleItemEventName
}

func GetCreateNoticeRuleItemEventNameEnum() CreateNoticeRuleItemEventNameEnum {
	return CreateNoticeRuleItemEventNameEnum{
		SCHEDULED: CreateNoticeRuleItemEventName{
			value: "Scheduled",
		},
		FAILED_SCHEDULING: CreateNoticeRuleItemEventName{
			value: "FailedScheduling",
		},
		HEALTHY: CreateNoticeRuleItemEventName{
			value: "Healthy",
		},
		UNHEALTHY: CreateNoticeRuleItemEventName{
			value: "Unhealthy",
		},
		PULLED: CreateNoticeRuleItemEventName{
			value: "Pulled",
		},
		FAILED_PULL_IMAGE: CreateNoticeRuleItemEventName{
			value: "FailedPullImage",
		},
		STARTED: CreateNoticeRuleItemEventName{
			value: "Started",
		},
		BACK_OFF_START: CreateNoticeRuleItemEventName{
			value: "BackOffStart",
		},
		SUCCESSFUL_MOUNT_VOLUME: CreateNoticeRuleItemEventName{
			value: "SuccessfulMountVolume",
		},
		FAILED_MOUNT: CreateNoticeRuleItemEventName{
			value: "FailedMount",
		},
	}
}

func (c CreateNoticeRuleItemEventName) Value() string {
	return c.value
}

func (c CreateNoticeRuleItemEventName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNoticeRuleItemEventName) UnmarshalJSON(b []byte) error {
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
