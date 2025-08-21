package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowNotificationSubscriptionResponse Response Object
type ShowNotificationSubscriptionResponse struct {

	// **参数解释：** 仓库ID。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 开启通知。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释：** 配资源。
	ConfigSource *ShowNotificationSubscriptionResponseConfigSource `json:"config_source,omitempty"`

	WebhookConfig *RepoWebHookSubscriptionDto `json:"webhook_config,omitempty"`

	// **参数解释：** 仓库使用量告警阀值（百分比）。
	WaringRepoUsageRate *ShowNotificationSubscriptionResponseWaringRepoUsageRate `json:"waring_repo_usage_rate,omitempty"`

	// **参数解释：** 通知事件。
	SubscriptEvents *[]RepoSubscriptionEventDto `json:"subscript_events,omitempty"`
	HttpStatusCode  int                         `json:"-"`
}

func (o ShowNotificationSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotificationSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"ShowNotificationSubscriptionResponse", string(data)}, " ")
}

type ShowNotificationSubscriptionResponseConfigSource struct {
	value string
}

type ShowNotificationSubscriptionResponseConfigSourceEnum struct {
	REPO    ShowNotificationSubscriptionResponseConfigSource
	PRODUCT ShowNotificationSubscriptionResponseConfigSource
}

func GetShowNotificationSubscriptionResponseConfigSourceEnum() ShowNotificationSubscriptionResponseConfigSourceEnum {
	return ShowNotificationSubscriptionResponseConfigSourceEnum{
		REPO: ShowNotificationSubscriptionResponseConfigSource{
			value: "repo",
		},
		PRODUCT: ShowNotificationSubscriptionResponseConfigSource{
			value: "product",
		},
	}
}

func (c ShowNotificationSubscriptionResponseConfigSource) Value() string {
	return c.value
}

func (c ShowNotificationSubscriptionResponseConfigSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowNotificationSubscriptionResponseConfigSource) UnmarshalJSON(b []byte) error {
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

type ShowNotificationSubscriptionResponseWaringRepoUsageRate struct {
	value int32
}

type ShowNotificationSubscriptionResponseWaringRepoUsageRateEnum struct {
	E_60 ShowNotificationSubscriptionResponseWaringRepoUsageRate
	E_80 ShowNotificationSubscriptionResponseWaringRepoUsageRate
	E_90 ShowNotificationSubscriptionResponseWaringRepoUsageRate
}

func GetShowNotificationSubscriptionResponseWaringRepoUsageRateEnum() ShowNotificationSubscriptionResponseWaringRepoUsageRateEnum {
	return ShowNotificationSubscriptionResponseWaringRepoUsageRateEnum{
		E_60: ShowNotificationSubscriptionResponseWaringRepoUsageRate{
			value: 60,
		}, E_80: ShowNotificationSubscriptionResponseWaringRepoUsageRate{
			value: 80,
		}, E_90: ShowNotificationSubscriptionResponseWaringRepoUsageRate{
			value: 90,
		},
	}
}

func (c ShowNotificationSubscriptionResponseWaringRepoUsageRate) Value() int32 {
	return c.value
}

func (c ShowNotificationSubscriptionResponseWaringRepoUsageRate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowNotificationSubscriptionResponseWaringRepoUsageRate) UnmarshalJSON(b []byte) error {
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
