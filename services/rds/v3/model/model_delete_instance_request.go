package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type DeleteInstanceRequest struct {
	// 语言

	XLanguage *DeleteInstanceRequestXLanguage `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o DeleteInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRequest", string(data)}, " ")
}

type DeleteInstanceRequestXLanguage struct {
	value string
}

type DeleteInstanceRequestXLanguageEnum struct {
	ZH_CN DeleteInstanceRequestXLanguage
	EN_US DeleteInstanceRequestXLanguage
}

func GetDeleteInstanceRequestXLanguageEnum() DeleteInstanceRequestXLanguageEnum {
	return DeleteInstanceRequestXLanguageEnum{
		ZH_CN: DeleteInstanceRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteInstanceRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteInstanceRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteInstanceRequestXLanguage) UnmarshalJSON(b []byte) error {
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
