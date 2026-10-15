package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceHealthReportConfigRequest Request Object
type ShowInstanceHealthReportConfigRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ShowInstanceHealthReportConfigRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowInstanceHealthReportConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceHealthReportConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceHealthReportConfigRequest", string(data)}, " ")
}

type ShowInstanceHealthReportConfigRequestXLanguage struct {
	value string
}

type ShowInstanceHealthReportConfigRequestXLanguageEnum struct {
	EN_US ShowInstanceHealthReportConfigRequestXLanguage
	ZH_CN ShowInstanceHealthReportConfigRequestXLanguage
}

func GetShowInstanceHealthReportConfigRequestXLanguageEnum() ShowInstanceHealthReportConfigRequestXLanguageEnum {
	return ShowInstanceHealthReportConfigRequestXLanguageEnum{
		EN_US: ShowInstanceHealthReportConfigRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowInstanceHealthReportConfigRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowInstanceHealthReportConfigRequestXLanguage) Value() string {
	return c.value
}

func (c ShowInstanceHealthReportConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceHealthReportConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
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
