package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowNotificationSubscriptionRequest Request Object
type ShowNotificationSubscriptionRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 通知类型。 **取值范围：**   - internal_message，站内信。   - email，邮件。   - qyweixin，企业微信。   - feishu，飞书。   - dingding，钉钉。
	Type ShowNotificationSubscriptionRequestType `json:"type"`
}

func (o ShowNotificationSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotificationSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ShowNotificationSubscriptionRequest", string(data)}, " ")
}

type ShowNotificationSubscriptionRequestType struct {
	value string
}

type ShowNotificationSubscriptionRequestTypeEnum struct {
	INTERNAL_MESSAGE ShowNotificationSubscriptionRequestType
	EMAIL            ShowNotificationSubscriptionRequestType
	QYWEIXIN         ShowNotificationSubscriptionRequestType
	FEISHU           ShowNotificationSubscriptionRequestType
	DINGDING         ShowNotificationSubscriptionRequestType
}

func GetShowNotificationSubscriptionRequestTypeEnum() ShowNotificationSubscriptionRequestTypeEnum {
	return ShowNotificationSubscriptionRequestTypeEnum{
		INTERNAL_MESSAGE: ShowNotificationSubscriptionRequestType{
			value: "internal_message",
		},
		EMAIL: ShowNotificationSubscriptionRequestType{
			value: "email",
		},
		QYWEIXIN: ShowNotificationSubscriptionRequestType{
			value: "qyweixin",
		},
		FEISHU: ShowNotificationSubscriptionRequestType{
			value: "feishu",
		},
		DINGDING: ShowNotificationSubscriptionRequestType{
			value: "dingding",
		},
	}
}

func (c ShowNotificationSubscriptionRequestType) Value() string {
	return c.value
}

func (c ShowNotificationSubscriptionRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowNotificationSubscriptionRequestType) UnmarshalJSON(b []byte) error {
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
