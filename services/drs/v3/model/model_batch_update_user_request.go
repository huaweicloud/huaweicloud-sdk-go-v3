package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchUpdateUserRequest Request Object
type BatchUpdateUserRequest struct {

	// 请求语言类型
	XLanguage *BatchUpdateUserRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchUpdateSrcUserReq `json:"body,omitempty"`
}

func (o BatchUpdateUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateUserRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateUserRequest", string(data)}, " ")
}

type BatchUpdateUserRequestXLanguage struct {
	value string
}

type BatchUpdateUserRequestXLanguageEnum struct {
	EN_US BatchUpdateUserRequestXLanguage
	ZH_CN BatchUpdateUserRequestXLanguage
}

func GetBatchUpdateUserRequestXLanguageEnum() BatchUpdateUserRequestXLanguageEnum {
	return BatchUpdateUserRequestXLanguageEnum{
		EN_US: BatchUpdateUserRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchUpdateUserRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchUpdateUserRequestXLanguage) Value() string {
	return c.value
}

func (c BatchUpdateUserRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateUserRequestXLanguage) UnmarshalJSON(b []byte) error {
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
