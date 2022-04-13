package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSecretsResponse struct {
	// 密钥

	Secrets *[]SecretDetailResp `json:"secrets,omitempty"`
	// 满足条件的个数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSecretsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretsResponse struct{}"
	}

	return strings.Join([]string{"ListSecretsResponse", string(data)}, " ")
}
