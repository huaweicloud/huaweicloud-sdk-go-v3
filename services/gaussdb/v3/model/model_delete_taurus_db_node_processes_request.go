package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteTaurusDbNodeProcessesRequest Request Object
type DeleteTaurusDbNodeProcessesRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us
	XLanguage *DeleteTaurusDbNodeProcessesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  请求体内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认值**：  不涉及。
	ContentType string `json:"Content-Type"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  节点ID，此参数是节点的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为no07，长度为36个字符。  **默认取值**：  不涉及。
	NodeId string `json:"node_id"`

	Body *DeleteTaurusDbNodeProcessesRequestBody `json:"body,omitempty"`
}

func (o DeleteTaurusDbNodeProcessesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaurusDbNodeProcessesRequest struct{}"
	}

	return strings.Join([]string{"DeleteTaurusDbNodeProcessesRequest", string(data)}, " ")
}

type DeleteTaurusDbNodeProcessesRequestXLanguage struct {
	value string
}

type DeleteTaurusDbNodeProcessesRequestXLanguageEnum struct {
	ZH_CN DeleteTaurusDbNodeProcessesRequestXLanguage
	EN_US DeleteTaurusDbNodeProcessesRequestXLanguage
}

func GetDeleteTaurusDbNodeProcessesRequestXLanguageEnum() DeleteTaurusDbNodeProcessesRequestXLanguageEnum {
	return DeleteTaurusDbNodeProcessesRequestXLanguageEnum{
		ZH_CN: DeleteTaurusDbNodeProcessesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteTaurusDbNodeProcessesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteTaurusDbNodeProcessesRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteTaurusDbNodeProcessesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteTaurusDbNodeProcessesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
