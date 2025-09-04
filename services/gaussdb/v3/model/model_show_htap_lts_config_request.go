package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHtapLtsConfigRequest Request Object
type ShowHtapLtsConfigRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage ShowHtapLtsConfigRequestXLanguage `json:"X-Language"`

	// **参数解释**：  HTAP标准版实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 企业project ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： HTAP标准版实例名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**： 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**： 查询记录数。  **约束限制**：  不涉及。  **取值范围**： 0-100。  **默认取值**： 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  索引位置，偏移量。从第一条数据偏移offset条数据后开始查询（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。  **约束限制**：  不涉及。  **取值范围**：     0-10000。  **默认取值**：  0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowHtapLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHtapLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowHtapLtsConfigRequest", string(data)}, " ")
}

type ShowHtapLtsConfigRequestXLanguage struct {
	value string
}

type ShowHtapLtsConfigRequestXLanguageEnum struct {
	ZH_CN ShowHtapLtsConfigRequestXLanguage
	EN_US ShowHtapLtsConfigRequestXLanguage
}

func GetShowHtapLtsConfigRequestXLanguageEnum() ShowHtapLtsConfigRequestXLanguageEnum {
	return ShowHtapLtsConfigRequestXLanguageEnum{
		ZH_CN: ShowHtapLtsConfigRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowHtapLtsConfigRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowHtapLtsConfigRequestXLanguage) Value() string {
	return c.value
}

func (c ShowHtapLtsConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHtapLtsConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
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
