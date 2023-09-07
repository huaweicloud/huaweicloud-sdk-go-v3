package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportOperationInfoRequest Request Object
type ExportOperationInfoRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ExportOperationInfoRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ExportOperationInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportOperationInfoRequest struct{}"
	}

	return strings.Join([]string{"ExportOperationInfoRequest", string(data)}, " ")
}

type ExportOperationInfoRequestXLanguage struct {
	value string
}

type ExportOperationInfoRequestXLanguageEnum struct {
	EN_US ExportOperationInfoRequestXLanguage
	ZH_CN ExportOperationInfoRequestXLanguage
}

func GetExportOperationInfoRequestXLanguageEnum() ExportOperationInfoRequestXLanguageEnum {
	return ExportOperationInfoRequestXLanguageEnum{
		EN_US: ExportOperationInfoRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExportOperationInfoRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExportOperationInfoRequestXLanguage) Value() string {
	return c.value
}

func (c ExportOperationInfoRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportOperationInfoRequestXLanguage) UnmarshalJSON(b []byte) error {
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
