package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchSetSmnRequest struct {

	// 请求语言类型
	XLanguage *BatchSetSmnRequestXLanguage `json:"X-Language,omitempty" xml:"X-Language"`

	Body *BatchImportSmnInfoReq `json:"body,omitempty" xml:"body"`
}

func (o BatchSetSmnRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSmnRequest struct{}"
	}

	return strings.Join([]string{"BatchSetSmnRequest", string(data)}, " ")
}

type BatchSetSmnRequestXLanguage struct {
	value string
}

type BatchSetSmnRequestXLanguageEnum struct {
	EN_US BatchSetSmnRequestXLanguage
	ZH_CN BatchSetSmnRequestXLanguage
}

func GetBatchSetSmnRequestXLanguageEnum() BatchSetSmnRequestXLanguageEnum {
	return BatchSetSmnRequestXLanguageEnum{
		EN_US: BatchSetSmnRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchSetSmnRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchSetSmnRequestXLanguage) Value() string {
	return c.value
}

func (c BatchSetSmnRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSetSmnRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
