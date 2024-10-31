package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHealthReportTaskRequest Request Object
type CreateHealthReportTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。
	XLanguage *CreateHealthReportTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateHealthReportReq `json:"body,omitempty"`
}

func (o CreateHealthReportTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthReportTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateHealthReportTaskRequest", string(data)}, " ")
}

type CreateHealthReportTaskRequestXLanguage struct {
	value string
}

type CreateHealthReportTaskRequestXLanguageEnum struct {
	EN_US CreateHealthReportTaskRequestXLanguage
	ZH_CN CreateHealthReportTaskRequestXLanguage
}

func GetCreateHealthReportTaskRequestXLanguageEnum() CreateHealthReportTaskRequestXLanguageEnum {
	return CreateHealthReportTaskRequestXLanguageEnum{
		EN_US: CreateHealthReportTaskRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateHealthReportTaskRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateHealthReportTaskRequestXLanguage) Value() string {
	return c.value
}

func (c CreateHealthReportTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHealthReportTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
