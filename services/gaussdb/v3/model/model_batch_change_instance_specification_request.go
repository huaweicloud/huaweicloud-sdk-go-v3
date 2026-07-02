package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchChangeInstanceSpecificationRequest Request Object
type BatchChangeInstanceSpecificationRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *BatchChangeInstanceSpecificationRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认值**：  application/json。
	ContentType string `json:"Content-Type"`

	Body *BatchChangeInstanceSpecificationRequestBody `json:"body,omitempty"`
}

func (o BatchChangeInstanceSpecificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeInstanceSpecificationRequest struct{}"
	}

	return strings.Join([]string{"BatchChangeInstanceSpecificationRequest", string(data)}, " ")
}

type BatchChangeInstanceSpecificationRequestXLanguage struct {
	value string
}

type BatchChangeInstanceSpecificationRequestXLanguageEnum struct {
	ZH_CN BatchChangeInstanceSpecificationRequestXLanguage
	EN_US BatchChangeInstanceSpecificationRequestXLanguage
}

func GetBatchChangeInstanceSpecificationRequestXLanguageEnum() BatchChangeInstanceSpecificationRequestXLanguageEnum {
	return BatchChangeInstanceSpecificationRequestXLanguageEnum{
		ZH_CN: BatchChangeInstanceSpecificationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: BatchChangeInstanceSpecificationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c BatchChangeInstanceSpecificationRequestXLanguage) Value() string {
	return c.value
}

func (c BatchChangeInstanceSpecificationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchChangeInstanceSpecificationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
