package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSessionOverviewRequest Request Object
type ShowSessionOverviewRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ShowSessionOverviewRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`
}

func (o ShowSessionOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSessionOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowSessionOverviewRequest", string(data)}, " ")
}

type ShowSessionOverviewRequestXLanguage struct {
	value string
}

type ShowSessionOverviewRequestXLanguageEnum struct {
	ZH_CN ShowSessionOverviewRequestXLanguage
	EN_US ShowSessionOverviewRequestXLanguage
}

func GetShowSessionOverviewRequestXLanguageEnum() ShowSessionOverviewRequestXLanguageEnum {
	return ShowSessionOverviewRequestXLanguageEnum{
		ZH_CN: ShowSessionOverviewRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSessionOverviewRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSessionOverviewRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSessionOverviewRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSessionOverviewRequestXLanguage) UnmarshalJSON(b []byte) error {
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
