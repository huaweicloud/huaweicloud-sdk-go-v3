package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchTransactionConfigResponse Response Object
type SearchTransactionConfigResponse struct {

	// URL跟踪视图配置列表。
	TransactionConfigItemList *[]TransactionConfigItem `json:"transaction_config_item_list,omitempty"`

	// 总页数。
	TotalPage *int32 `json:"total_page,omitempty"`

	// 总配置数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SearchTransactionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchTransactionConfigResponse struct{}"
	}

	return strings.Join([]string{"SearchTransactionConfigResponse", string(data)}, " ")
}
