package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateBackupResourcePackageRequest Request Object
type CreateBackupResourcePackageRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us
	XLanguage *CreateBackupResourcePackageRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认值**：  application/json。
	ContentType string `json:"Content-Type"`

	Body *CreateBackupResourcePackageRequestBody `json:"body,omitempty"`
}

func (o CreateBackupResourcePackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackupResourcePackageRequest struct{}"
	}

	return strings.Join([]string{"CreateBackupResourcePackageRequest", string(data)}, " ")
}

type CreateBackupResourcePackageRequestXLanguage struct {
	value string
}

type CreateBackupResourcePackageRequestXLanguageEnum struct {
	ZH_CN CreateBackupResourcePackageRequestXLanguage
	EN_US CreateBackupResourcePackageRequestXLanguage
}

func GetCreateBackupResourcePackageRequestXLanguageEnum() CreateBackupResourcePackageRequestXLanguageEnum {
	return CreateBackupResourcePackageRequestXLanguageEnum{
		ZH_CN: CreateBackupResourcePackageRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateBackupResourcePackageRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateBackupResourcePackageRequestXLanguage) Value() string {
	return c.value
}

func (c CreateBackupResourcePackageRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBackupResourcePackageRequestXLanguage) UnmarshalJSON(b []byte) error {
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
