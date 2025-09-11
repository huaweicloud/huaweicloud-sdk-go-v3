package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListKmsTdeKeyRequest Request Object
type ListKmsTdeKeyRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListKmsTdeKeyRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: GaussDB使用透明加密的KMS主密钥ID所在资源空间名称。 获取方法请参见[获取项目名称](https://support.huaweicloud.com/api-gaussdb/gaussdb_api_196.html)。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	KmsProjectName string `json:"kms_project_name"`
}

func (o ListKmsTdeKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKmsTdeKeyRequest struct{}"
	}

	return strings.Join([]string{"ListKmsTdeKeyRequest", string(data)}, " ")
}

type ListKmsTdeKeyRequestXLanguage struct {
	value string
}

type ListKmsTdeKeyRequestXLanguageEnum struct {
	ZH_CN ListKmsTdeKeyRequestXLanguage
	EN_US ListKmsTdeKeyRequestXLanguage
}

func GetListKmsTdeKeyRequestXLanguageEnum() ListKmsTdeKeyRequestXLanguageEnum {
	return ListKmsTdeKeyRequestXLanguageEnum{
		ZH_CN: ListKmsTdeKeyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListKmsTdeKeyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListKmsTdeKeyRequestXLanguage) Value() string {
	return c.value
}

func (c ListKmsTdeKeyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListKmsTdeKeyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
