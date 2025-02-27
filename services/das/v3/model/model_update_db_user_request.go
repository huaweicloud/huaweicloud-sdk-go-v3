package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateDbUserRequest Request Object
type UpdateDbUserRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库用户ID。每个用户需绑定一个数据库账号（数据库用户由“注册数据库用户”接口创建）。
	DbUserId string `json:"db_user_id"`

	// 语言
	XLanguage *UpdateDbUserRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdateDbUserRequestBody `json:"body,omitempty"`
}

func (o UpdateDbUserRequest) String() string {
	data, err := utils.Marshal(o)
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

func (c UpdateDbUserRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateDbUserRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDbUserRequestXLanguage) UnmarshalJSON(b []byte) error {
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
