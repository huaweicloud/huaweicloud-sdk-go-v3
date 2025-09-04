package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceNodeRequest Request Object
type ListInstanceNodeRequest struct {

	// **参数解释**：  HTAP标准版实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage ListInstanceNodeRequestXLanguage `json:"X-Language"`
}

func (o ListInstanceNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceNodeRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceNodeRequest", string(data)}, " ")
}

type ListInstanceNodeRequestXLanguage struct {
	value string
}

type ListInstanceNodeRequestXLanguageEnum struct {
	ZH_CN ListInstanceNodeRequestXLanguage
	EN_US ListInstanceNodeRequestXLanguage
}

func GetListInstanceNodeRequestXLanguageEnum() ListInstanceNodeRequestXLanguageEnum {
	return ListInstanceNodeRequestXLanguageEnum{
		ZH_CN: ListInstanceNodeRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListInstanceNodeRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListInstanceNodeRequestXLanguage) Value() string {
	return c.value
}

func (c ListInstanceNodeRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceNodeRequestXLanguage) UnmarshalJSON(b []byte) error {
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
