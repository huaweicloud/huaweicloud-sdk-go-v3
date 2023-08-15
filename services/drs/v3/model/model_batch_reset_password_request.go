package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchResetPasswordRequest Request Object
type BatchResetPasswordRequest struct {

	// 请求语言类型
	XLanguage *BatchResetPasswordRequestXLanguage `json:"X-Language,omitempty"`

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

func (c BatchResetPasswordRequestXLanguage) Value() string {
	return c.value
}

func (c BatchResetPasswordRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchResetPasswordRequestXLanguage) UnmarshalJSON(b []byte) error {
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
