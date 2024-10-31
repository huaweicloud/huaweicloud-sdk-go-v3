package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceHealthReportRequest Request Object
type ShowInstanceHealthReportRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 任务ID。
	TaskId string `json:"task_id"`

	// 请求语言类型。
	XLanguage *ShowInstanceHealthReportRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowInstanceHealthReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceHealthReportRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceHealthReportRequest", string(data)}, " ")
}

type ShowInstanceHealthReportRequestXLanguage struct {
	value string
}

type ShowInstanceHealthReportRequestXLanguageEnum struct {
	EN_US ShowInstanceHealthReportRequestXLanguage
	ZH_CN ShowInstanceHealthReportRequestXLanguage
}

func GetShowInstanceHealthReportRequestXLanguageEnum() ShowInstanceHealthReportRequestXLanguageEnum {
	return ShowInstanceHealthReportRequestXLanguageEnum{
		EN_US: ShowInstanceHealthReportRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowInstanceHealthReportRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowInstanceHealthReportRequestXLanguage) Value() string {
	return c.value
}

func (c ShowInstanceHealthReportRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceHealthReportRequestXLanguage) UnmarshalJSON(b []byte) error {
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
