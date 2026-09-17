package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryTransactionsResponse Response Object
type ListHistoryTransactionsResponse struct {

	// 历史事务总数
	Total *int32 `json:"total,omitempty"`

	// 历史事务信息列表
	TransactionInfoList *[]TransactionInfo `json:"transaction_info_list,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o ListHistoryTransactionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryTransactionsResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryTransactionsResponse", string(data)}, " ")
}
