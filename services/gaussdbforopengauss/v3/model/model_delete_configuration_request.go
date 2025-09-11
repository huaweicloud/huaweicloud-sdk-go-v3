package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteConfigurationRequest Request Object
type DeleteConfigurationRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *DeleteConfigurationRequestXLanguage `json:"X-Language,omitempty"`

	// 参数配置模板ID。
	ConfigId string `json:"config_id"`
}

func (o DeleteConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigurationRequest struct{}"
	}

	return strings.Join([]string{"DeleteConfigurationRequest", string(data)}, " ")
}

type DeleteConfigurationRequestXLanguage struct {
	value string
}

type DeleteConfigurationRequestXLanguageEnum struct {
	ZH_CN DeleteConfigurationRequestXLanguage
	EN_US DeleteConfigurationRequestXLanguage
}

func GetDeleteConfigurationRequestXLanguageEnum() DeleteConfigurationRequestXLanguageEnum {
	return DeleteConfigurationRequestXLanguageEnum{
		ZH_CN: DeleteConfigurationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteConfigurationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteConfigurationRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteConfigurationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteConfigurationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
