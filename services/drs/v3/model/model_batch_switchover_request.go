package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchSwitchoverRequest struct {
	// 请求语言类型

	XLanguage BatchSwitchoverRequestXLanguage `json:"X-Language"`

	Body *BatchSwitchoverReq `json:"body,omitempty"`
}

func (o BatchSwitchoverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSwitchoverRequest struct{}"
	}

	return strings.Join([]string{"BatchSwitchoverRequest", string(data)}, " ")
}

type BatchSwitchoverRequestXLanguage struct {
	value string
}

type BatchSwitchoverRequestXLanguageEnum struct {
	EN_US BatchSwitchoverRequestXLanguage
	ZH_CN BatchSwitchoverRequestXLanguage
}

func GetBatchSwitchoverRequestXLanguageEnum() BatchSwitchoverRequestXLanguageEnum {
	return BatchSwitchoverRequestXLanguageEnum{
		EN_US: BatchSwitchoverRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchSwitchoverRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchSwitchoverRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSwitchoverRequestXLanguage) UnmarshalJSON(b []byte) error {
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
