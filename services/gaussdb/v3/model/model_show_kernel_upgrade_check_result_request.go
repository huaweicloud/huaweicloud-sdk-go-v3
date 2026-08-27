package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowKernelUpgradeCheckResultRequest Request Object
type ShowKernelUpgradeCheckResultRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *ShowKernelUpgradeCheckResultRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpgradePrecheckRequest `json:"body,omitempty"`
}

func (o ShowKernelUpgradeCheckResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKernelUpgradeCheckResultRequest struct{}"
	}

	return strings.Join([]string{"ShowKernelUpgradeCheckResultRequest", string(data)}, " ")
}

type ShowKernelUpgradeCheckResultRequestXLanguage struct {
	value string
}

type ShowKernelUpgradeCheckResultRequestXLanguageEnum struct {
	ZH_CN ShowKernelUpgradeCheckResultRequestXLanguage
	EN_US ShowKernelUpgradeCheckResultRequestXLanguage
}

func GetShowKernelUpgradeCheckResultRequestXLanguageEnum() ShowKernelUpgradeCheckResultRequestXLanguageEnum {
	return ShowKernelUpgradeCheckResultRequestXLanguageEnum{
		ZH_CN: ShowKernelUpgradeCheckResultRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowKernelUpgradeCheckResultRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowKernelUpgradeCheckResultRequestXLanguage) Value() string {
	return c.value
}

func (c ShowKernelUpgradeCheckResultRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowKernelUpgradeCheckResultRequestXLanguage) UnmarshalJSON(b []byte) error {
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
