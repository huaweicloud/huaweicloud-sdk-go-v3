package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVaultResponse Response Object
type ListVaultResponse struct {

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

func (o ListVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVaultResponse struct{}"
	}

	return strings.Join([]string{"ListVaultResponse", string(data)}, " ")
}
