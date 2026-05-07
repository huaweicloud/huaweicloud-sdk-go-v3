package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListUpgradePathsRequest Request Object
type ListUpgradePathsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListUpgradePathsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 源引擎版本号。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	SourceVersion string `json:"source_version"`

	// **参数解释**: 目标引擎版本号。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	TargetVersion *string `json:"target_version,omitempty"`
}

func (o ListUpgradePathsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpgradePathsRequest struct{}"
	}

	return strings.Join([]string{"ListUpgradePathsRequest", string(data)}, " ")
}

type ListUpgradePathsRequestXLanguage struct {
	value string
}

type ListUpgradePathsRequestXLanguageEnum struct {
	ZH_CN ListUpgradePathsRequestXLanguage
	EN_US ListUpgradePathsRequestXLanguage
}

func GetListUpgradePathsRequestXLanguageEnum() ListUpgradePathsRequestXLanguageEnum {
	return ListUpgradePathsRequestXLanguageEnum{
		ZH_CN: ListUpgradePathsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListUpgradePathsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListUpgradePathsRequestXLanguage) Value() string {
	return c.value
}

func (c ListUpgradePathsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUpgradePathsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
