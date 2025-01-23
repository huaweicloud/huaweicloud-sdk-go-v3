package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMetricNamesSupportRequest Request Object
type ShowMetricNamesSupportRequest struct {

	// 语言
	XLanguage *ShowMetricNamesSupportRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowMetricNamesSupportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricNamesSupportRequest struct{}"
	}

	return strings.Join([]string{"ShowMetricNamesSupportRequest", string(data)}, " ")
}

type ShowMetricNamesSupportRequestXLanguage struct {
	value string
}

type ShowMetricNamesSupportRequestXLanguageEnum struct {
	ZH_CN ShowMetricNamesSupportRequestXLanguage
	EN_US ShowMetricNamesSupportRequestXLanguage
}

func GetShowMetricNamesSupportRequestXLanguageEnum() ShowMetricNamesSupportRequestXLanguageEnum {
	return ShowMetricNamesSupportRequestXLanguageEnum{
		ZH_CN: ShowMetricNamesSupportRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowMetricNamesSupportRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowMetricNamesSupportRequestXLanguage) Value() string {
	return c.value
}

func (c ShowMetricNamesSupportRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetricNamesSupportRequestXLanguage) UnmarshalJSON(b []byte) error {
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
