package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateDynamicServerlessPolicyRequest Request Object
type UpdateDynamicServerlessPolicyRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us
	XLanguage *UpdateDynamicServerlessPolicyRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认值**：  application/json。
	ContentType string `json:"Content-Type"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	Body *UpdateDynamicServerlessPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateDynamicServerlessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDynamicServerlessPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDynamicServerlessPolicyRequest", string(data)}, " ")
}

type UpdateDynamicServerlessPolicyRequestXLanguage struct {
	value string
}

type UpdateDynamicServerlessPolicyRequestXLanguageEnum struct {
	ZH_CN UpdateDynamicServerlessPolicyRequestXLanguage
	EN_US UpdateDynamicServerlessPolicyRequestXLanguage
}

func GetUpdateDynamicServerlessPolicyRequestXLanguageEnum() UpdateDynamicServerlessPolicyRequestXLanguageEnum {
	return UpdateDynamicServerlessPolicyRequestXLanguageEnum{
		ZH_CN: UpdateDynamicServerlessPolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateDynamicServerlessPolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateDynamicServerlessPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateDynamicServerlessPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDynamicServerlessPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
