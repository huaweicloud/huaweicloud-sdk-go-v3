package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSecretsResponse struct {

	// 数量
	Count *int32 `json:"count,omitempty"`

	// 缪瑶详情列表
	Secrets        *[]SecretDetail `json:"secrets,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListSecretsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretsResponse struct{}"
	}

	return strings.Join([]string{"ListSecretsResponse", string(data)}, " ")
}
