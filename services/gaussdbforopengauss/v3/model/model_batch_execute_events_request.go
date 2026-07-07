package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchExecuteEventsRequest Request Object
type BatchExecuteEventsRequest struct {

	// **参数解释**: 请求语言类型。 **约束限制**: 不涉及。 **取值范围**: - en-us - zh-cn  **默认取值**: en-us。
	XLanguage *BatchExecuteEventsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchExecuteEventsRequestBody `json:"body,omitempty"`
}

func (o BatchExecuteEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExecuteEventsRequest struct{}"
	}

	return strings.Join([]string{"BatchExecuteEventsRequest", string(data)}, " ")
}

type BatchExecuteEventsRequestXLanguage struct {
	value string
}

type BatchExecuteEventsRequestXLanguageEnum struct {
	EN_US BatchExecuteEventsRequestXLanguage
	ZH_CN BatchExecuteEventsRequestXLanguage
}

func GetBatchExecuteEventsRequestXLanguageEnum() BatchExecuteEventsRequestXLanguageEnum {
	return BatchExecuteEventsRequestXLanguageEnum{
		EN_US: BatchExecuteEventsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchExecuteEventsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchExecuteEventsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchExecuteEventsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchExecuteEventsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
