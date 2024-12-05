package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteManualBackupRequest Request Object
type BatchDeleteManualBackupRequest struct {

	// 语言。
	XLanguage *BatchDeleteManualBackupRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchDeleteManualBackupRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteManualBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteManualBackupRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteManualBackupRequest", string(data)}, " ")
}

type BatchDeleteManualBackupRequestXLanguage struct {
	value string
}

type BatchDeleteManualBackupRequestXLanguageEnum struct {
	ZH_CN BatchDeleteManualBackupRequestXLanguage
	EN_US BatchDeleteManualBackupRequestXLanguage
}

func GetBatchDeleteManualBackupRequestXLanguageEnum() BatchDeleteManualBackupRequestXLanguageEnum {
	return BatchDeleteManualBackupRequestXLanguageEnum{
		ZH_CN: BatchDeleteManualBackupRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: BatchDeleteManualBackupRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c BatchDeleteManualBackupRequestXLanguage) Value() string {
	return c.value
}

func (c BatchDeleteManualBackupRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteManualBackupRequestXLanguage) UnmarshalJSON(b []byte) error {
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
