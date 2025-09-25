package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHbaInfoHistoryResponse Response Object
type ListHbaInfoHistoryResponse struct {

	// **参数解释**: hba修改历史信息。
	HbaHistories *[]HbaHistoryResult `json:"hba_histories,omitempty"`

	// **参数解释**: hba配置总数。 **取值范围**: 不涉及。
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListHbaInfoHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHbaInfoHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListHbaInfoHistoryResponse", string(data)}, " ")
}
