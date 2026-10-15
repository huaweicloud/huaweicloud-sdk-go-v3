package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCheckJobsRequest Request Object
type BatchCheckJobsRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文 - zh-cn：中文  **默认取值**：  en-us
	XLanguage *BatchCheckJobsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchPrecheckReq `json:"body,omitempty"`
}

func (o BatchCheckJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchCheckJobsRequest", string(data)}, " ")
}

type BatchCheckJobsRequestXLanguage struct {
	value string
}

type BatchCheckJobsRequestXLanguageEnum struct {
	EN_US BatchCheckJobsRequestXLanguage
	ZH_CN BatchCheckJobsRequestXLanguage
}

func GetBatchCheckJobsRequestXLanguageEnum() BatchCheckJobsRequestXLanguageEnum {
	return BatchCheckJobsRequestXLanguageEnum{
		EN_US: BatchCheckJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchCheckJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchCheckJobsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchCheckJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCheckJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
