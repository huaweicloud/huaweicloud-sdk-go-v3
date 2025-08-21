package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateRepoNotificationSubscriptionDto struct {

	// **参数解释：** 开启通知。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释：** 配资源。
	ConfigSource *UpdateRepoNotificationSubscriptionDtoConfigSource `json:"config_source,omitempty"`

	WebhookConfig *UpdateRepoWebHookSubscriptionDto `json:"webhook_config,omitempty"`

	// **参数解释：** 仓库使用量告警阀值（百分比）。
	WaringRepoUsageRate *UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRate `json:"waring_repo_usage_rate,omitempty"`

	// **参数解释：** 通知事件。
	SubscriptEvents *[]RepoSubscriptionEventDto `json:"subscript_events,omitempty"`
}

func (o UpdateRepoNotificationSubscriptionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepoNotificationSubscriptionDto struct{}"
	}

	return strings.Join([]string{"UpdateRepoNotificationSubscriptionDto", string(data)}, " ")
}

type UpdateRepoNotificationSubscriptionDtoConfigSource struct {
	value string
}

type UpdateRepoNotificationSubscriptionDtoConfigSourceEnum struct {
	REPO    UpdateRepoNotificationSubscriptionDtoConfigSource
	PRODUCT UpdateRepoNotificationSubscriptionDtoConfigSource
}

func GetUpdateRepoNotificationSubscriptionDtoConfigSourceEnum() UpdateRepoNotificationSubscriptionDtoConfigSourceEnum {
	return UpdateRepoNotificationSubscriptionDtoConfigSourceEnum{
		REPO: UpdateRepoNotificationSubscriptionDtoConfigSource{
			value: "repo",
		},
		PRODUCT: UpdateRepoNotificationSubscriptionDtoConfigSource{
			value: "product",
		},
	}
}

func (c UpdateRepoNotificationSubscriptionDtoConfigSource) Value() string {
	return c.value
}

func (c UpdateRepoNotificationSubscriptionDtoConfigSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRepoNotificationSubscriptionDtoConfigSource) UnmarshalJSON(b []byte) error {
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

type UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRate struct {
	value int32
}

type UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRateEnum struct {
	E_60 UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRate
	E_80 UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRate
	E_90 UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRate
}

func GetUpdateRepoNotificationSubscriptionDtoWaringRepoUsageRateEnum() UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRateEnum {
	return UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRateEnum{
		E_60: UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRate{
			value: 60,
		}, E_80: UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRate{
			value: 80,
		}, E_90: UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRate{
			value: 90,
		},
	}
}

func (c UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRate) Value() int32 {
	return c.value
}

func (c UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRepoNotificationSubscriptionDtoWaringRepoUsageRate) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
