package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchUpdateUserRequest struct {
	// 请求语言类型

	XLanguage BatchUpdateUserRequestXLanguage `json:"X-Language"`

	Body *BatchUpdateSrcUserReq `json:"body,omitempty"`
}

func (o BatchUpdateUserRequest) String() string {
	data, err := json.Marshal(o)
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

func (c BatchUpdateUserRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchUpdateUserRequestXLanguage) UnmarshalJSON(b []byte) error {
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
