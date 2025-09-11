package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowKmsKeyDetailRequest Request Object
type ShowKmsKeyDetailRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ShowKmsKeyDetailRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: kms秘钥ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	KmsKeyId string `json:"kms_key_id"`

	// **参数解释**: GaussDB使用透明加密的KMS主密钥ID所在资源空间名称。 获取方法请参见[获取项目名称](https://support.huaweicloud.com/api-gaussdb/gaussdb_api_196.html)。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	KmsProjectName string `json:"kms_project_name"`
}

func (o ShowKmsKeyDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKmsKeyDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowKmsKeyDetailRequest", string(data)}, " ")
}

type ShowKmsKeyDetailRequestXLanguage struct {
	value string
}

type ShowKmsKeyDetailRequestXLanguageEnum struct {
	ZH_CN ShowKmsKeyDetailRequestXLanguage
	EN_US ShowKmsKeyDetailRequestXLanguage
}

func GetShowKmsKeyDetailRequestXLanguageEnum() ShowKmsKeyDetailRequestXLanguageEnum {
	return ShowKmsKeyDetailRequestXLanguageEnum{
		ZH_CN: ShowKmsKeyDetailRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowKmsKeyDetailRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowKmsKeyDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowKmsKeyDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowKmsKeyDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
