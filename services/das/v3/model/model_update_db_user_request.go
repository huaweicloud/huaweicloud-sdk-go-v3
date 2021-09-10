package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdateDbUserRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 数据库用户ID

	DbUserId string `json:"db_user_id"`
	// 语言

	XLanguage *UpdateDbUserRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdateDbUserRequestBody `json:"body,omitempty"`
}

func (o UpdateDbUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDbUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateDbUserRequest", string(data)}, " ")
}

type UpdateDbUserRequestXLanguage struct {
	value string
}

type UpdateDbUserRequestXLanguageEnum struct {
	ZH_CN UpdateDbUserRequestXLanguage
	EN_US UpdateDbUserRequestXLanguage
}

func GetUpdateDbUserRequestXLanguageEnum() UpdateDbUserRequestXLanguageEnum {
	return UpdateDbUserRequestXLanguageEnum{
		ZH_CN: UpdateDbUserRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateDbUserRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateDbUserRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateDbUserRequestXLanguage) UnmarshalJSON(b []byte) error {
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
