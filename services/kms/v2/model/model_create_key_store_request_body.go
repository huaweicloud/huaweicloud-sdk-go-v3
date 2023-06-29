package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKeyStoreRequestBody 创建专属密钥库请求体
type CreateKeyStoreRequestBody struct {

	// 专属密钥库别名，取值范围为1到255个字符，满足正则匹配“^[a-zA-Z0-9:/_-]{1,255}$”，且不与已有的专属密钥库别名重名。
	KeystoreAlias string `json:"keystore_alias"`

	// DHSM集群Id，要求集群当前未创建专属密钥库。
	HsmClusterId string `json:"hsm_cluster_id"`

	// DHSM集群的CA证书
	HsmCaCert string `json:"hsm_ca_cert"`
}

func (o CreateKeyStoreRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyStoreRequestBody struct{}"
	}

	return strings.Join([]string{"CreateKeyStoreRequestBody", string(data)}, " ")
}
