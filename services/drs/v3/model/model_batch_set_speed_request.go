package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchSetSpeedRequest struct {

	// 请求语言类型
	XLanguage *BatchSetSpeedRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchLimitSpeedReq `json:"body,omitempty"`
}

func (o BatchSetSpeedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSpeedRequest struct{}"
	}

	return strings.Join([]string{"BatchSetSpeedRequest", string(data)}, " ")
}

type BatchSetSpeedRequestXLanguage struct {
	value string
}

type BatchSetSpeedRequestXLanguageEnum struct {
	EN_US BatchSetSpeedRequestXLanguage
	ZH_CN BatchSetSpeedRequestXLanguage
}

func GetBatchSetSpeedRequestXLanguageEnum() BatchSetSpeedRequestXLanguageEnum {
	return BatchSetSpeedRequestXLanguageEnum{
		EN_US: BatchSetSpeedRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchSetSpeedRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchSetSpeedRequestXLanguage) Value() string {
	return c.value
}

func (c BatchSetSpeedRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSetSpeedRequestXLanguage) UnmarshalJSON(b []byte) error {
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
