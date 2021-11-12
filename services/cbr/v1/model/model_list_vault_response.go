package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListVaultResponse struct {
	// 存储库实例列表

	Vaults *[]Vault `json:"vaults,omitempty"`
	// 存储库个数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVaultResponse struct{}"
	}

	return strings.Join([]string{"ListVaultResponse", string(data)}, " ")
}
