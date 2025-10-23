package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVaultTopsResponse Response Object
type ListVaultTopsResponse struct {

	// 存储库总数
	Count *int64 `json:"count,omitempty"`

	// 存储库统计信息
	Vaults         *[]ListVaultTopsResponseBodyVaults `json:"vaults,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListVaultTopsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVaultTopsResponse struct{}"
	}

	return strings.Join([]string{"ListVaultTopsResponse", string(data)}, " ")
}
