package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchListRposAndRtosRequest Request Object
type BatchListRposAndRtosRequest struct {

	// 请求语言类型
	XLanguage *BatchListRposAndRtosRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchQueryRpoAndRtoReq `json:"body,omitempty"`
}

func (o BatchListRposAndRtosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListRposAndRtosRequest struct{}"
	}

	return strings.Join([]string{"BatchListRposAndRtosRequest", string(data)}, " ")
}

type BatchListRposAndRtosRequestXLanguage struct {
	value string
}

type BatchListRposAndRtosRequestXLanguageEnum struct {
	EN_US BatchListRposAndRtosRequestXLanguage
	ZH_CN BatchListRposAndRtosRequestXLanguage
}

func GetBatchListRposAndRtosRequestXLanguageEnum() BatchListRposAndRtosRequestXLanguageEnum {
	return BatchListRposAndRtosRequestXLanguageEnum{
		EN_US: BatchListRposAndRtosRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchListRposAndRtosRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchListRposAndRtosRequestXLanguage) Value() string {
	return c.value
}

func (c BatchListRposAndRtosRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListRposAndRtosRequestXLanguage) UnmarshalJSON(b []byte) error {
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
