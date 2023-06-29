package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchTransactionResponse Response Object
type SearchTransactionResponse struct {

	// URL跟踪视图列表。
	TxItemList *[]TxItemVo `json:"tx_item_list,omitempty"`

	// 最后响应时间。
	LatestTime *int64 `json:"latest_time,omitempty"`

	// 总数据条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 请求id。
	ResultId       *string `json:"result_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SearchTransactionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchTransactionResponse struct{}"
	}

	return strings.Join([]string{"SearchTransactionResponse", string(data)}, " ")
}
