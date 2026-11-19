package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplyConfigurationWithEntitiesRequest Request Object
type ApplyConfigurationWithEntitiesRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ApplyConfigurationWithEntitiesRequestXLanguage `json:"X-Language,omitempty"`

	// 参数模板ID。
	ConfigId string `json:"config_id"`

	Body *ApplyConfigurationEntityRequestBody `json:"body,omitempty"`
}

func (o ApplyConfigurationWithEntitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationWithEntitiesRequest struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationWithEntitiesRequest", string(data)}, " ")
}

type ApplyConfigurationWithEntitiesRequestXLanguage struct {
	value string
}

type ApplyConfigurationWithEntitiesRequestXLanguageEnum struct {
	ZH_CN ApplyConfigurationWithEntitiesRequestXLanguage
	EN_US ApplyConfigurationWithEntitiesRequestXLanguage
}

func GetApplyConfigurationWithEntitiesRequestXLanguageEnum() ApplyConfigurationWithEntitiesRequestXLanguageEnum {
	return ApplyConfigurationWithEntitiesRequestXLanguageEnum{
		ZH_CN: ApplyConfigurationWithEntitiesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ApplyConfigurationWithEntitiesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ApplyConfigurationWithEntitiesRequestXLanguage) Value() string {
	return c.value
}

func (c ApplyConfigurationWithEntitiesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplyConfigurationWithEntitiesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
