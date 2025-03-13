package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecretsResponse Response Object
type ListSecretsResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// 资源种类。
	Kind *string `json:"kind,omitempty"`

	// 定时启停规则所包含的所有应用，只在生效范围为application的时候需要填写。
	Items          *[]SecretDetail `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListSecretsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretsResponse struct{}"
	}

	return strings.Join([]string{"ListSecretsResponse", string(data)}, " ")
}
