package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateNotificationSubscriptionResponse Response Object
type UpdateNotificationSubscriptionResponse struct {

	// **参数解释：** 仓库ID。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 开启通知。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释：** 配资源。
	ConfigSource *UpdateNotificationSubscriptionResponseConfigSource `json:"config_source,omitempty"`

	WebhookConfig *RepoWebHookSubscriptionDto `json:"webhook_config,omitempty"`

	// **参数解释：** 仓库使用量告警阀值（百分比）。
	WaringRepoUsageRate *UpdateNotificationSubscriptionResponseWaringRepoUsageRate `json:"waring_repo_usage_rate,omitempty"`

	// **参数解释：** 通知事件。
	SubscriptEvents *[]RepoSubscriptionEventDto `json:"subscript_events,omitempty"`
	HttpStatusCode  int                         `json:"-"`
}

func (o UpdateNotificationSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"UpdateNotificationSubscriptionResponse", string(data)}, " ")
}

type UpdateNotificationSubscriptionResponseConfigSource struct {
	value string
}

type UpdateNotificationSubscriptionResponseConfigSourceEnum struct {
	REPO    UpdateNotificationSubscriptionResponseConfigSource
	PRODUCT UpdateNotificationSubscriptionResponseConfigSource
}

func GetUpdateNotificationSubscriptionResponseConfigSourceEnum() UpdateNotificationSubscriptionResponseConfigSourceEnum {
	return UpdateNotificationSubscriptionResponseConfigSourceEnum{
		REPO: UpdateNotificationSubscriptionResponseConfigSource{
			value: "repo",
		},
		PRODUCT: UpdateNotificationSubscriptionResponseConfigSource{
			value: "product",
		},
	}
}

func (c UpdateNotificationSubscriptionResponseConfigSource) Value() string {
	return c.value
}

func (c UpdateNotificationSubscriptionResponseConfigSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationSubscriptionResponseConfigSource) UnmarshalJSON(b []byte) error {
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

type UpdateNotificationSubscriptionResponseWaringRepoUsageRate struct {
	value int32
}

type UpdateNotificationSubscriptionResponseWaringRepoUsageRateEnum struct {
	E_60 UpdateNotificationSubscriptionResponseWaringRepoUsageRate
	E_80 UpdateNotificationSubscriptionResponseWaringRepoUsageRate
	E_90 UpdateNotificationSubscriptionResponseWaringRepoUsageRate
}

func GetUpdateNotificationSubscriptionResponseWaringRepoUsageRateEnum() UpdateNotificationSubscriptionResponseWaringRepoUsageRateEnum {
	return UpdateNotificationSubscriptionResponseWaringRepoUsageRateEnum{
		E_60: UpdateNotificationSubscriptionResponseWaringRepoUsageRate{
			value: 60,
		}, E_80: UpdateNotificationSubscriptionResponseWaringRepoUsageRate{
			value: 80,
		}, E_90: UpdateNotificationSubscriptionResponseWaringRepoUsageRate{
			value: 90,
		},
	}
}

func (c UpdateNotificationSubscriptionResponseWaringRepoUsageRate) Value() int32 {
	return c.value
}

func (c UpdateNotificationSubscriptionResponseWaringRepoUsageRate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationSubscriptionResponseWaringRepoUsageRate) UnmarshalJSON(b []byte) error {
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
