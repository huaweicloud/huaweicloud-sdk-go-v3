package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpgradeConfigurationRequest Request Object
type UpgradeConfigurationRequest struct {

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认取值**：  application/json。
	ContentType string `json:"Content-Type"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *UpgradeConfigurationRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  参数模板ID，此参数是参数模板的唯一标识。 获取方法请参见[查询参数模板](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlConfigurations.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	ConfigurationId string `json:"configuration_id"`

	Body *UpgradeConfigurationRequestBody `json:"body,omitempty"`
}

func (o UpgradeConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpgradeConfigurationRequest", string(data)}, " ")
}

type UpgradeConfigurationRequestXLanguage struct {
	value string
}

type UpgradeConfigurationRequestXLanguageEnum struct {
	ZH_CN UpgradeConfigurationRequestXLanguage
	EN_US UpgradeConfigurationRequestXLanguage
}

func GetUpgradeConfigurationRequestXLanguageEnum() UpgradeConfigurationRequestXLanguageEnum {
	return UpgradeConfigurationRequestXLanguageEnum{
		ZH_CN: UpgradeConfigurationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpgradeConfigurationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpgradeConfigurationRequestXLanguage) Value() string {
	return c.value
}

func (c UpgradeConfigurationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpgradeConfigurationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
