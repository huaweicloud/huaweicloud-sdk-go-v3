package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchListJobDetailsRequest struct {

	// 请求语言类型
	XLanguage *BatchListJobDetailsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchQueryJobReqPage `json:"body,omitempty"`
}

func (o BatchListJobDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListJobDetailsRequest struct{}"
	}

	return strings.Join([]string{"BatchListJobDetailsRequest", string(data)}, " ")
}

type BatchListJobDetailsRequestXLanguage struct {
	value string
}

type BatchListJobDetailsRequestXLanguageEnum struct {
	EN_US BatchListJobDetailsRequestXLanguage
	ZH_CN BatchListJobDetailsRequestXLanguage
}

func GetBatchListJobDetailsRequestXLanguageEnum() BatchListJobDetailsRequestXLanguageEnum {
	return BatchListJobDetailsRequestXLanguageEnum{
		EN_US: BatchListJobDetailsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchListJobDetailsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchListJobDetailsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchListJobDetailsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListJobDetailsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
