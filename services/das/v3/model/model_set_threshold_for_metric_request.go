package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SetThresholdForMetricRequest Request Object
type SetThresholdForMetricRequest struct {

	// 语言
	XLanguage *SetThresholdForMetricRequestXLanguage `json:"X-Language,omitempty"`

	Body *ApiSetMetricCodeThresholdReq `json:"body,omitempty"`
}

func (o SetThresholdForMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetThresholdForMetricRequest struct{}"
	}

	return strings.Join([]string{"SetThresholdForMetricRequest", string(data)}, " ")
}

type SetThresholdForMetricRequestXLanguage struct {
	value string
}

type SetThresholdForMetricRequestXLanguageEnum struct {
	ZH_CN SetThresholdForMetricRequestXLanguage
	EN_US SetThresholdForMetricRequestXLanguage
}

func GetSetThresholdForMetricRequestXLanguageEnum() SetThresholdForMetricRequestXLanguageEnum {
	return SetThresholdForMetricRequestXLanguageEnum{
		ZH_CN: SetThresholdForMetricRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SetThresholdForMetricRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SetThresholdForMetricRequestXLanguage) Value() string {
	return c.value
}

func (c SetThresholdForMetricRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetThresholdForMetricRequestXLanguage) UnmarshalJSON(b []byte) error {
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
