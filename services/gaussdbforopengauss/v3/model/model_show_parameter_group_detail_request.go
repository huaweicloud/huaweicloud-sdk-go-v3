package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowParameterGroupDetailRequest Request Object
type ShowParameterGroupDetailRequest struct {

	// 语言,默认：en-us。
	XLanguage *ShowParameterGroupDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 参数模板ID
	ConfigId string `json:"config_id"`
}

func (o ShowParameterGroupDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowParameterGroupDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowParameterGroupDetailRequest", string(data)}, " ")
}

type ShowParameterGroupDetailRequestXLanguage struct {
	value string
}

type ShowParameterGroupDetailRequestXLanguageEnum struct {
	ZH_CN ShowParameterGroupDetailRequestXLanguage
	EN_US ShowParameterGroupDetailRequestXLanguage
}

func GetShowParameterGroupDetailRequestXLanguageEnum() ShowParameterGroupDetailRequestXLanguageEnum {
	return ShowParameterGroupDetailRequestXLanguageEnum{
		ZH_CN: ShowParameterGroupDetailRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowParameterGroupDetailRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowParameterGroupDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowParameterGroupDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowParameterGroupDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
