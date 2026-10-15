package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SetInstanceHealthReportConfigRequest Request Object
type SetInstanceHealthReportConfigRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *SetInstanceHealthReportConfigRequestXLanguage `json:"X-Language,omitempty"`

	Body *SetInstanceHealthReportConfigReq `json:"body,omitempty"`
}

func (o SetInstanceHealthReportConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInstanceHealthReportConfigRequest struct{}"
	}

	return strings.Join([]string{"SetInstanceHealthReportConfigRequest", string(data)}, " ")
}

type SetInstanceHealthReportConfigRequestXLanguage struct {
	value string
}

type SetInstanceHealthReportConfigRequestXLanguageEnum struct {
	EN_US SetInstanceHealthReportConfigRequestXLanguage
	ZH_CN SetInstanceHealthReportConfigRequestXLanguage
}

func GetSetInstanceHealthReportConfigRequestXLanguageEnum() SetInstanceHealthReportConfigRequestXLanguageEnum {
	return SetInstanceHealthReportConfigRequestXLanguageEnum{
		EN_US: SetInstanceHealthReportConfigRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: SetInstanceHealthReportConfigRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c SetInstanceHealthReportConfigRequestXLanguage) Value() string {
	return c.value
}

func (c SetInstanceHealthReportConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetInstanceHealthReportConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
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
