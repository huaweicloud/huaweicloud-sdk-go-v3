package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteDbUserRequest Request Object
type DeleteDbUserRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库用户ID。每个用户需绑定一个数据库账号（数据库用户由“注册数据库用户”接口创建）。
	DbUserId string `json:"db_user_id"`

	// 语言
	XLanguage *DeleteDbUserRequestXLanguage `json:"X-Language,omitempty"`
}

func (o DeleteDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteDbUserRequest", string(data)}, " ")
}

type DeleteDbUserRequestXLanguage struct {
	value string
}

type DeleteDbUserRequestXLanguageEnum struct {
	ZH_CN DeleteDbUserRequestXLanguage
	EN_US DeleteDbUserRequestXLanguage
}

func GetDeleteDbUserRequestXLanguageEnum() DeleteDbUserRequestXLanguageEnum {
	return DeleteDbUserRequestXLanguageEnum{
		ZH_CN: DeleteDbUserRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteDbUserRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteDbUserRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteDbUserRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDbUserRequestXLanguage) UnmarshalJSON(b []byte) error {
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
