package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DeleteProcessRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 语言

	XLanguage *DeleteProcessRequestXLanguage `json:"X-Language,omitempty"`

	Body *DeleteProcessReqBody `json:"body,omitempty"`
}

func (o DeleteProcessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProcessRequest struct{}"
	}

	return strings.Join([]string{"DeleteProcessRequest", string(data)}, " ")
}

type DeleteProcessRequestXLanguage struct {
	value string
}

type DeleteProcessRequestXLanguageEnum struct {
	ZH_CN DeleteProcessRequestXLanguage
	EN_US DeleteProcessRequestXLanguage
}

func GetDeleteProcessRequestXLanguageEnum() DeleteProcessRequestXLanguageEnum {
	return DeleteProcessRequestXLanguageEnum{
		ZH_CN: DeleteProcessRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteProcessRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteProcessRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteProcessRequestXLanguage) UnmarshalJSON(b []byte) error {
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
