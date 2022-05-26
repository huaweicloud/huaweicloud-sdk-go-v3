package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchRestoreTaskRequest struct {

	// 请求语言类型
	XLanguage *BatchRestoreTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchRetryReq `json:"body,omitempty"`
}

func (o BatchRestoreTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestoreTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchRestoreTaskRequest", string(data)}, " ")
}

type BatchRestoreTaskRequestXLanguage struct {
	value string
}

type BatchRestoreTaskRequestXLanguageEnum struct {
	EN_US BatchRestoreTaskRequestXLanguage
	ZH_CN BatchRestoreTaskRequestXLanguage
}

func GetBatchRestoreTaskRequestXLanguageEnum() BatchRestoreTaskRequestXLanguageEnum {
	return BatchRestoreTaskRequestXLanguageEnum{
		EN_US: BatchRestoreTaskRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchRestoreTaskRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchRestoreTaskRequestXLanguage) Value() string {
	return c.value
}

func (c BatchRestoreTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRestoreTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
