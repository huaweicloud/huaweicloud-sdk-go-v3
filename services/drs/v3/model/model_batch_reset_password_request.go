package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type BatchResetPasswordRequest struct {
	// 请求语言类型

	XLanguage BatchResetPasswordRequestXLanguage `json:"X-Language"`

	Body *BatchModifyPwdReq `json:"body,omitempty"`
}

func (o BatchResetPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResetPasswordRequest struct{}"
	}

	return strings.Join([]string{"BatchResetPasswordRequest", string(data)}, " ")
}

type BatchResetPasswordRequestXLanguage struct {
	value string
}

type BatchResetPasswordRequestXLanguageEnum struct {
	EN_US BatchResetPasswordRequestXLanguage
	ZH_CN BatchResetPasswordRequestXLanguage
}

func GetBatchResetPasswordRequestXLanguageEnum() BatchResetPasswordRequestXLanguageEnum {
	return BatchResetPasswordRequestXLanguageEnum{
		EN_US: BatchResetPasswordRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchResetPasswordRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchResetPasswordRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchResetPasswordRequestXLanguage) UnmarshalJSON(b []byte) error {
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
