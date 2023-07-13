package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoredValueCardsResponse Response Object
type ListStoredValueCardsResponse struct {

	// 符合查询条件的总条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 优惠券记录。 具体请参见表2。
	StoredValueCards *[]UserStoredValueCard `json:"stored_value_cards,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o ListStoredValueCardsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoredValueCardsResponse struct{}"
	}

	return strings.Join([]string{"ListStoredValueCardsResponse", string(data)}, " ")
}
