package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RestartInstanceRequest Request Object
type RestartInstanceRequest struct {

	// 语言
	XLanguage *RestartInstanceRequestXLanguage `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o RestartInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestartInstanceRequest", string(data)}, " ")
}

type RestartInstanceRequestXLanguage struct {
	value string
}

type RestartInstanceRequestXLanguageEnum struct {
	ZH_CN RestartInstanceRequestXLanguage
	EN_US RestartInstanceRequestXLanguage
}

func GetRestartInstanceRequestXLanguageEnum() RestartInstanceRequestXLanguageEnum {
	return RestartInstanceRequestXLanguageEnum{
		ZH_CN: RestartInstanceRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: RestartInstanceRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c RestartInstanceRequestXLanguage) Value() string {
	return c.value
}

func (c RestartInstanceRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestartInstanceRequestXLanguage) UnmarshalJSON(b []byte) error {
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
