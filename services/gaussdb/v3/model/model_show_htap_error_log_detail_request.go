package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHtapErrorLogDetailRequest Request Object
type ShowHtapErrorLogDetailRequest struct {

	// **参数解释**：  HTAP标准版实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage ShowHtapErrorLogDetailRequestXLanguage `json:"X-Language"`

	Body *HtapErrorLogQueryRequestBody `json:"body,omitempty"`
}

func (o ShowHtapErrorLogDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHtapErrorLogDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowHtapErrorLogDetailRequest", string(data)}, " ")
}

type ShowHtapErrorLogDetailRequestXLanguage struct {
	value string
}

type ShowHtapErrorLogDetailRequestXLanguageEnum struct {
	ZH_CN ShowHtapErrorLogDetailRequestXLanguage
	EN_US ShowHtapErrorLogDetailRequestXLanguage
}

func GetShowHtapErrorLogDetailRequestXLanguageEnum() ShowHtapErrorLogDetailRequestXLanguageEnum {
	return ShowHtapErrorLogDetailRequestXLanguageEnum{
		ZH_CN: ShowHtapErrorLogDetailRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowHtapErrorLogDetailRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowHtapErrorLogDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowHtapErrorLogDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHtapErrorLogDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
