package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateNoticeRuleItem struct {

	// 触发事件名称，支持实例调度成功、实例调度失败、健康检查成功、健康检查失败、镜像拉取成功、镜像拉取失败、容器启动成功、容器启动失败、卷挂载成功、卷挂载失败。
	EventName *UpdateNoticeRuleItemEventName `json:"event_name,omitempty"`

	Scope *NoticeRuleScope `json:"scope"`

	TriggerPolicy *TriggerPolicy `json:"trigger_policy"`

	// 是否启用。
	Enable *bool `json:"enable,omitempty"`
}

func (o UpdateNoticeRuleItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNoticeRuleItem struct{}"
	}

	return strings.Join([]string{"UpdateNoticeRuleItem", string(data)}, " ")
}

type UpdateNoticeRuleItemEventName struct {
	value string
}

type UpdateNoticeRuleItemEventNameEnum struct {
	SCHEDULED               UpdateNoticeRuleItemEventName
	FAILED_SCHEDULING       UpdateNoticeRuleItemEventName
	HEALTHY                 UpdateNoticeRuleItemEventName
	UNHEALTHY               UpdateNoticeRuleItemEventName
	PULLED                  UpdateNoticeRuleItemEventName
	FAILED_PULL_IMAGE       UpdateNoticeRuleItemEventName
	STARTED                 UpdateNoticeRuleItemEventName
	BACK_OFF_START          UpdateNoticeRuleItemEventName
	SUCCESSFUL_MOUNT_VOLUME UpdateNoticeRuleItemEventName
	FAILED_MOUNT            UpdateNoticeRuleItemEventName
}

func GetUpdateNoticeRuleItemEventNameEnum() UpdateNoticeRuleItemEventNameEnum {
	return UpdateNoticeRuleItemEventNameEnum{
		SCHEDULED: UpdateNoticeRuleItemEventName{
			value: "Scheduled",
		},
		FAILED_SCHEDULING: UpdateNoticeRuleItemEventName{
			value: "FailedScheduling",
		},
		HEALTHY: UpdateNoticeRuleItemEventName{
			value: "Healthy",
		},
		UNHEALTHY: UpdateNoticeRuleItemEventName{
			value: "Unhealthy",
		},
		PULLED: UpdateNoticeRuleItemEventName{
			value: "Pulled",
		},
		FAILED_PULL_IMAGE: UpdateNoticeRuleItemEventName{
			value: "FailedPullImage",
		},
		STARTED: UpdateNoticeRuleItemEventName{
			value: "Started",
		},
		BACK_OFF_START: UpdateNoticeRuleItemEventName{
			value: "BackOffStart",
		},
		SUCCESSFUL_MOUNT_VOLUME: UpdateNoticeRuleItemEventName{
			value: "SuccessfulMountVolume",
		},
		FAILED_MOUNT: UpdateNoticeRuleItemEventName{
			value: "FailedMount",
		},
	}
}

func (c UpdateNoticeRuleItemEventName) Value() string {
	return c.value
}

func (c UpdateNoticeRuleItemEventName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNoticeRuleItemEventName) UnmarshalJSON(b []byte) error {
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
