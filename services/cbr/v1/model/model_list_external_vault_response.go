package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExternalVaultResponse Response Object
type ListExternalVaultResponse struct {

	// 存储库实例列表
	Vaults *[]Vault `json:"vaults,omitempty"`

	// 存储库个数
	Count *int32 `json:"count,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询
	Offset         *int32 `json:"offset,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListExternalVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExternalVaultResponse struct{}"
	}

	return strings.Join([]string{"ListExternalVaultResponse", string(data)}, " ")
}
