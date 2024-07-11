package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransactionsResponse Response Object
type ListTransactionsResponse struct {

	// 历史事务总数
	Total *int32 `json:"total,omitempty"`

	// 历史事务信息列表
	TransactionInfoList *[]GetTransactionListRespTransactionInfoList `json:"transaction_info_list,omitempty"`
	HttpStatusCode      int                                          `json:"-"`
}

func (o ListTransactionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransactionsResponse struct{}"
	}

	return strings.Join([]string{"ListTransactionsResponse", string(data)}, " ")
}
