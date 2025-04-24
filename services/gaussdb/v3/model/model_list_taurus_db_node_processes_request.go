package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTaurusDbNodeProcessesRequest Request Object
type ListTaurusDbNodeProcessesRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us
	XLanguage *ListTaurusDbNodeProcessesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  节点ID，此参数是节点的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为no07，长度为36个字符。  **默认取值**：  不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**：  索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  ≥0  **默认取值**：  0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  查询记录数。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  1-100  **默认取值**：  100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTaurusDbNodeProcessesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaurusDbNodeProcessesRequest struct{}"
	}

	return strings.Join([]string{"ListTaurusDbNodeProcessesRequest", string(data)}, " ")
}

type ListTaurusDbNodeProcessesRequestXLanguage struct {
	value string
}

type ListTaurusDbNodeProcessesRequestXLanguageEnum struct {
	ZH_CN ListTaurusDbNodeProcessesRequestXLanguage
	EN_US ListTaurusDbNodeProcessesRequestXLanguage
}

func GetListTaurusDbNodeProcessesRequestXLanguageEnum() ListTaurusDbNodeProcessesRequestXLanguageEnum {
	return ListTaurusDbNodeProcessesRequestXLanguageEnum{
		ZH_CN: ListTaurusDbNodeProcessesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTaurusDbNodeProcessesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTaurusDbNodeProcessesRequestXLanguage) Value() string {
	return c.value
}

func (c ListTaurusDbNodeProcessesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTaurusDbNodeProcessesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
