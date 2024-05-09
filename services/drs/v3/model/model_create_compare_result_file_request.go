package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateCompareResultFileRequest Request Object
type CreateCompareResultFileRequest struct {

	// 请求语言类型。
	XLanguage *CreateCompareResultFileRequestXLanguage `json:"X-Language,omitempty"`

	// 区域ID，例如：cn-north-4。
	Region string `json:"Region"`

	// 任务ID。
	JobId string `json:"job_id"`

	Body *ExportCompareResultReq `json:"body,omitempty"`
}

func (o CreateCompareResultFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompareResultFileRequest struct{}"
	}

	return strings.Join([]string{"CreateCompareResultFileRequest", string(data)}, " ")
}

type CreateCompareResultFileRequestXLanguage struct {
	value string
}

type CreateCompareResultFileRequestXLanguageEnum struct {
	EN_US CreateCompareResultFileRequestXLanguage
	ZH_CN CreateCompareResultFileRequestXLanguage
}

func GetCreateCompareResultFileRequestXLanguageEnum() CreateCompareResultFileRequestXLanguageEnum {
	return CreateCompareResultFileRequestXLanguageEnum{
		EN_US: CreateCompareResultFileRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateCompareResultFileRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateCompareResultFileRequestXLanguage) Value() string {
	return c.value
}

func (c CreateCompareResultFileRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCompareResultFileRequestXLanguage) UnmarshalJSON(b []byte) error {
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
