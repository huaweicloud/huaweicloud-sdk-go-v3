package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretVersionRequestBody 创建凭据版本请求体。
type CreateSecretVersionRequestBody struct {

	// 新创建凭据的凭据值，将其加密后，存入初始版本中。  类型：base64编码的二进制数据对象。  约束：secret_binary和secret_string必须且只能设置一个，最大32K。
	SecretBinary *string `json:"secret_binary,omitempty"`

	// 新创建凭据的凭据值，将其加密后，存入初始版本中。  约束：secret_binary和 secret_string必须且只能设置一个，最大32K。
	SecretString *string `json:"secret_string,omitempty"`

	// 凭据版本在存入时需要被同时标记的版本状态。如果您不指定此参数，凭据管家默认为新版本标记SYSCURRENT  约束：数组大小：最小1，最大12。stage长度：最小1字节，最大64字节。
	VersionStages *[]string `json:"version_stages,omitempty"`
}

func (o CreateSecretVersionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretVersionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecretVersionRequestBody", string(data)}, " ")
}
