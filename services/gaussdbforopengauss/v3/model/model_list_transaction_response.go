package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransactionResponse Response Object
type ListTransactionResponse struct {

	// **参数解释**: 查到的事务数量。 **取值范围**: 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: 事务信息列表。
	Rows           *[]ListTransactionResponseBodyRows `json:"rows,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListTransactionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransactionResponse struct{}"
	}

	return strings.Join([]string{"ListTransactionResponse", string(data)}, " ")
}
