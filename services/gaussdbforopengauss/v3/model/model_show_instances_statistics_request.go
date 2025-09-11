package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstancesStatisticsRequest Request Object
type ShowInstancesStatisticsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ShowInstancesStatisticsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowInstancesStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstancesStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowInstancesStatisticsRequest", string(data)}, " ")
}

type ShowInstancesStatisticsRequestXLanguage struct {
	value string
}

type ShowInstancesStatisticsRequestXLanguageEnum struct {
	ZH_CN ShowInstancesStatisticsRequestXLanguage
	EN_US ShowInstancesStatisticsRequestXLanguage
}

func GetShowInstancesStatisticsRequestXLanguageEnum() ShowInstancesStatisticsRequestXLanguageEnum {
	return ShowInstancesStatisticsRequestXLanguageEnum{
		ZH_CN: ShowInstancesStatisticsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowInstancesStatisticsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowInstancesStatisticsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowInstancesStatisticsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstancesStatisticsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
