package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteBackupRequest Request Object
type BatchDeleteBackupRequest struct {

	// 语言。
	XLanguage *BatchDeleteBackupRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchDeleteBackupRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBackupRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteBackupRequest", string(data)}, " ")
}

type BatchDeleteBackupRequestXLanguage struct {
	value string
}

type BatchDeleteBackupRequestXLanguageEnum struct {
	ZH_CN BatchDeleteBackupRequestXLanguage
	EN_US BatchDeleteBackupRequestXLanguage
}

func GetBatchDeleteBackupRequestXLanguageEnum() BatchDeleteBackupRequestXLanguageEnum {
	return BatchDeleteBackupRequestXLanguageEnum{
		ZH_CN: BatchDeleteBackupRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: BatchDeleteBackupRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c BatchDeleteBackupRequestXLanguage) Value() string {
	return c.value
}

func (c BatchDeleteBackupRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteBackupRequestXLanguage) UnmarshalJSON(b []byte) error {
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
