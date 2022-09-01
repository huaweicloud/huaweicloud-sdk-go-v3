package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DeleteDbUserRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 数据库用户ID
	DbUserId string `json:"db_user_id" xml:"db_user_id"`

	// 语言
	XLanguage *DeleteDbUserRequestXLanguage `json:"X-Language,omitempty" xml:"X-Language"`
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
