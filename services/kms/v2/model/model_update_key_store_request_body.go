package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeyStoreRequestBody 更新专属密钥库请求体
type UpdateKeyStoreRequestBody struct {

	// 访问XKS-Proxy的凭据ID
	ProxyAccessKeyId string `json:"proxy_access_key_id"`

	// 访问XKS-Proxy的凭据Key
	ProxySecretAccessKey string `json:"proxy_secret_access_key"`
}

func (o UpdateKeyStoreRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyStoreRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKeyStoreRequestBody", string(data)}, " ")
}
