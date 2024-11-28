package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EncryptionConfig struct {

	// **参数解释**： 加密模式，取值为Default或KMS。默认为Default，即使用cce本地密钥加密。若使用KMS加密模式则使用用户自定义密钥或默认密钥加密secret资源。 若用户在创建时未填写，则集群查询接口中默认会返回Default。 **约束限制**： 不涉及 **取值范围**： 取值范围：[Default, KMS]; **默认取值**： Default
	Mode *EncryptionConfigMode `json:"mode,omitempty"`

	// **参数解释**： kms密钥ID - 集群创建API中，如果mode字段设置为Default，无需填写该字段；如果mode字段设置为KMS，则支持填写该字段。若字段为空，则默认使用KMS默认密钥进行填充，默认密钥不存在时云服务将自动为用户创建cce/default默认密钥。 用户需使用真实存在的KMS密钥，并且在集群生命周期结束前，禁止删除、禁用密钥等操作，防止集群功能异常（集群设置该密钥后不允许修改）。  - 集群查询API中，如果mode字段设置为Default，则该字段返回为空；若mode字段设置为KMS，则该字段为具体的密钥ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	KmsKeyID *string `json:"kmsKeyID,omitempty"`
}

func (o EncryptionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptionConfig struct{}"
	}

	return strings.Join([]string{"EncryptionConfig", string(data)}, " ")
}

type EncryptionConfigMode struct {
	value string
}

type EncryptionConfigModeEnum struct {
	DEFAULT EncryptionConfigMode
	KMS     EncryptionConfigMode
}

func GetEncryptionConfigModeEnum() EncryptionConfigModeEnum {
	return EncryptionConfigModeEnum{
		DEFAULT: EncryptionConfigMode{
			value: "Default",
		},
		KMS: EncryptionConfigMode{
			value: "KMS",
		},
	}
}

func (c EncryptionConfigMode) Value() string {
	return c.value
}

func (c EncryptionConfigMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EncryptionConfigMode) UnmarshalJSON(b []byte) error {
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
