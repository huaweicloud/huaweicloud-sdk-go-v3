package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RepoNotificationSubscriptionStateDto struct {

	// **参数解释：** 配资源。
	ConfigSource *RepoNotificationSubscriptionStateDtoConfigSource `json:"config_source,omitempty"`

	// **参数解释：** 开启通知。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o RepoNotificationSubscriptionStateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoNotificationSubscriptionStateDto struct{}"
	}

	return strings.Join([]string{"RepoNotificationSubscriptionStateDto", string(data)}, " ")
}

type RepoNotificationSubscriptionStateDtoConfigSource struct {
	value string
}

type RepoNotificationSubscriptionStateDtoConfigSourceEnum struct {
	REPO    RepoNotificationSubscriptionStateDtoConfigSource
	PRODUCT RepoNotificationSubscriptionStateDtoConfigSource
}

func GetRepoNotificationSubscriptionStateDtoConfigSourceEnum() RepoNotificationSubscriptionStateDtoConfigSourceEnum {
	return RepoNotificationSubscriptionStateDtoConfigSourceEnum{
		REPO: RepoNotificationSubscriptionStateDtoConfigSource{
			value: "repo",
		},
		PRODUCT: RepoNotificationSubscriptionStateDtoConfigSource{
			value: "product",
		},
	}
}

func (c RepoNotificationSubscriptionStateDtoConfigSource) Value() string {
	return c.value
}

func (c RepoNotificationSubscriptionStateDtoConfigSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepoNotificationSubscriptionStateDtoConfigSource) UnmarshalJSON(b []byte) error {
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
