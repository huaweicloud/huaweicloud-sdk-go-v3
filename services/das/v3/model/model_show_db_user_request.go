package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowDbUserRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库用户ID
	DbUserId string `json:"db_user_id"`

	// 语言
	XLanguage *ShowDbUserRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbUserRequest struct{}"
	}

	return strings.Join([]string{"ShowDbUserRequest", string(data)}, " ")
}

type ShowDbUserRequestXLanguage struct {
	value string
}

type ShowDbUserRequestXLanguageEnum struct {
	ZH_CN ShowDbUserRequestXLanguage
	EN_US ShowDbUserRequestXLanguage
}

func GetShowDbUserRequestXLanguageEnum() ShowDbUserRequestXLanguageEnum {
	return ShowDbUserRequestXLanguageEnum{
		ZH_CN: ShowDbUserRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowDbUserRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowDbUserRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDbUserRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDbUserRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
