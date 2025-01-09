package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCsrRequest Request Object
type ListCsrRequest struct {

	// 每页条目数量，取值如下： - 10：每页显示10条证书信息。 - 20：每页显示20条证书信息。 - 50：每页显示50条证书信息。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// csr名称。
	Name *string `json:"name,omitempty"`

	// 密钥算法的类型。取值如下： - RSA_2048 - RSA_3072 - RSA_4096 - EC_P256 - EC_P384 - SM2
	PrivateKeyAlgo *string `json:"private_key_algo,omitempty"`
}

func (o ListCsrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCsrRequest struct{}"
	}

	return strings.Join([]string{"ListCsrRequest", string(data)}, " ")
}
