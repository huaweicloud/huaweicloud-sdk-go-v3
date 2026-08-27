package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckKernelUpgradeRequest Request Object
type CheckKernelUpgradeRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *CheckKernelUpgradeRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpgradePrecheckRequest `json:"body,omitempty"`
}

func (o CheckKernelUpgradeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckKernelUpgradeRequest struct{}"
	}

	return strings.Join([]string{"CheckKernelUpgradeRequest", string(data)}, " ")
}

type CheckKernelUpgradeRequestXLanguage struct {
	value string
}

type CheckKernelUpgradeRequestXLanguageEnum struct {
	ZH_CN CheckKernelUpgradeRequestXLanguage
	EN_US CheckKernelUpgradeRequestXLanguage
}

func GetCheckKernelUpgradeRequestXLanguageEnum() CheckKernelUpgradeRequestXLanguageEnum {
	return CheckKernelUpgradeRequestXLanguageEnum{
		ZH_CN: CheckKernelUpgradeRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CheckKernelUpgradeRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CheckKernelUpgradeRequestXLanguage) Value() string {
	return c.value
}

func (c CheckKernelUpgradeRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckKernelUpgradeRequestXLanguage) UnmarshalJSON(b []byte) error {
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
