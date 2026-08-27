package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaurusDbTxnProgressResponse Response Object
type ShowTaurusDbTxnProgressResponse struct {

	// **参数解释**：  处于活跃状态（回滚中）的事务进度列表。如果输入的ID已结束或不存在，则不在此列表中返回。
	Transactions *[]TxnItem `json:"transactions,omitempty"`

	// **参数解释**： 满足查询条件的事务记录总数。 **取值范围**： 0~100。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTaurusDbTxnProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaurusDbTxnProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowTaurusDbTxnProgressResponse", string(data)}, " ")
}
