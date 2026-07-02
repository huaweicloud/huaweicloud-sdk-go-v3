package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowBackupResourcePackageFlavorsRequest Request Object
type ShowBackupResourcePackageFlavorsRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us
	XLanguage *ShowBackupResourcePackageFlavorsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  ≥0。  **默认取值**：  0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  查询记录数。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  1-100。  **默认取值**：  100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowBackupResourcePackageFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupResourcePackageFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupResourcePackageFlavorsRequest", string(data)}, " ")
}

type ShowBackupResourcePackageFlavorsRequestXLanguage struct {
	value string
}

type ShowBackupResourcePackageFlavorsRequestXLanguageEnum struct {
	ZH_CN ShowBackupResourcePackageFlavorsRequestXLanguage
	EN_US ShowBackupResourcePackageFlavorsRequestXLanguage
}

func GetShowBackupResourcePackageFlavorsRequestXLanguageEnum() ShowBackupResourcePackageFlavorsRequestXLanguageEnum {
	return ShowBackupResourcePackageFlavorsRequestXLanguageEnum{
		ZH_CN: ShowBackupResourcePackageFlavorsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowBackupResourcePackageFlavorsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowBackupResourcePackageFlavorsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowBackupResourcePackageFlavorsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBackupResourcePackageFlavorsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
