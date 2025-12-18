package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateNotificationSubscriptionRequest Request Object
type UpdateNotificationSubscriptionRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 通知类型。 **取值范围：**   - internal_message，站内信。   - email，邮件。   - qyweixin，企业微信。   - feishu，飞书。   - dingding，钉钉。
	Type UpdateNotificationSubscriptionRequestType `json:"type"`

	Body *UpdateRepoNotificationSubscriptionDto `json:"body,omitempty"`
}

func (o UpdateNotificationSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateNotificationSubscriptionRequest", string(data)}, " ")
}

type UpdateNotificationSubscriptionRequestType struct {
	value string
}

type UpdateNotificationSubscriptionRequestTypeEnum struct {
	INTERNAL_MESSAGE UpdateNotificationSubscriptionRequestType
	EMAIL            UpdateNotificationSubscriptionRequestType
	QYWEIXIN         UpdateNotificationSubscriptionRequestType
	FEISHU           UpdateNotificationSubscriptionRequestType
	DINGDING         UpdateNotificationSubscriptionRequestType
}

func GetUpdateNotificationSubscriptionRequestTypeEnum() UpdateNotificationSubscriptionRequestTypeEnum {
	return UpdateNotificationSubscriptionRequestTypeEnum{
		INTERNAL_MESSAGE: UpdateNotificationSubscriptionRequestType{
			value: "internal_message",
		},
		EMAIL: UpdateNotificationSubscriptionRequestType{
			value: "email",
		},
		QYWEIXIN: UpdateNotificationSubscriptionRequestType{
			value: "qyweixin",
		},
		FEISHU: UpdateNotificationSubscriptionRequestType{
			value: "feishu",
		},
		DINGDING: UpdateNotificationSubscriptionRequestType{
			value: "dingding",
		},
	}
}

func (c UpdateNotificationSubscriptionRequestType) Value() string {
	return c.value
}

func (c UpdateNotificationSubscriptionRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationSubscriptionRequestType) UnmarshalJSON(b []byte) error {
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
