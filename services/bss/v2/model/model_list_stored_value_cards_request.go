package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoredValueCardsRequest Request Object
type ListStoredValueCardsRequest struct {

	// 状态：1：可使用2：已用完
	Status int32 `json:"status"`

	// 储值卡ID。此参数不携带或携带值为空时，不作为筛选条件。
	CardId *string `json:"card_id,omitempty"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的储值卡的数量，默认值为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListStoredValueCardsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoredValueCardsRequest struct{}"
	}

	return strings.Join([]string{"ListStoredValueCardsRequest", string(data)}, " ")
}
