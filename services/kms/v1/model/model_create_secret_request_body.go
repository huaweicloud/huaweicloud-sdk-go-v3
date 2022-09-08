package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建凭据请求消息体。
type CreateSecretRequestBody struct {

	// 凭据名称。  约束：取值范围为1到64个字符，满足正则匹配“^[a-zA-Z0-9._-]{1,64}$”。
	Name string `json:"name"`

	// 用于加密保护凭据值的KMS主密钥ID，如果您未指定此参数，凭据管理服务将默认使用名为csms/default的默认主密钥，用于加密您账号在本项目中创建的凭据值。如果用户账号下不存在该名称的主密钥，则凭据管理服务自动为您创建该名称的密钥。
	KmsKeyId *string `json:"kms_key_id,omitempty"`

	// 凭据的描述信息。  约束：2048字节。
	Description *string `json:"description,omitempty"`

	// 二进制类型凭据在base64编码后的明文，凭据管理服务将其加密后，存入凭据的初始版本中。  类型：base64编码的二进制数据对象。  约束：secret_binary和secret_string必须且只能设置一个，最大32K。
	SecretBinary *string `json:"secret_binary,omitempty"`

	// 文本类型凭据的明文，凭据管理服务将其加密后，存入凭据的初始版本中。  约束：secret_binary和secret_string必须且只能设置一个，最大32K。
	SecretString *string `json:"secret_string,omitempty"`
}

func (o CreateSecretRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecretRequestBody", string(data)}, " ")
}
