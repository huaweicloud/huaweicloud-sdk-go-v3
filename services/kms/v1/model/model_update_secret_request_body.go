package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新凭据对象的元数据信息请求体
type UpdateSecretRequestBody struct {
	// 凭据名称。  约束：取值范围为1到64个字符，满足正则匹配“^[a-zA-Z0-9._-]{1,64}$”。

	Name *string `json:"name,omitempty"`
	// 用于加密保护凭据值的KMS主密钥ID。更新凭据的主密钥后，仅新创建的凭据版本使用更新后的主密钥ID加密，之前的凭据版本依旧使用之前的主密钥ID解密。

	KmsKeyId *string `json:"kms_key_id,omitempty"`
	// 凭据的描述信息。  约束：2048字节。

	Description *string `json:"description,omitempty"`
}

func (o UpdateSecretRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecretRequestBody", string(data)}, " ")
}
