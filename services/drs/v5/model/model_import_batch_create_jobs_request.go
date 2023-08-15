package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImportBatchCreateJobsRequest Request Object
type ImportBatchCreateJobsRequest struct {

	// 请求语言类型。
	XLanguage *ImportBatchCreateJobsRequestXLanguage `json:"X-Language,omitempty"`

	Body *ImportBatchCreateJobsRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportBatchCreateJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportBatchCreateJobsRequest struct{}"
	}

	return strings.Join([]string{"ImportBatchCreateJobsRequest", string(data)}, " ")
}

type ImportBatchCreateJobsRequestXLanguage struct {
	value string
}

type ImportBatchCreateJobsRequestXLanguageEnum struct {
	EN_US ImportBatchCreateJobsRequestXLanguage
	ZH_CN ImportBatchCreateJobsRequestXLanguage
}

func GetImportBatchCreateJobsRequestXLanguageEnum() ImportBatchCreateJobsRequestXLanguageEnum {
	return ImportBatchCreateJobsRequestXLanguageEnum{
		EN_US: ImportBatchCreateJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ImportBatchCreateJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ImportBatchCreateJobsRequestXLanguage) Value() string {
	return c.value
}

func (c ImportBatchCreateJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportBatchCreateJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
