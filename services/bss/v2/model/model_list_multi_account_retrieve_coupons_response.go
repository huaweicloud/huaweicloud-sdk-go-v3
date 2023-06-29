package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMultiAccountRetrieveCouponsResponse Response Object
type ListMultiAccountRetrieveCouponsResponse struct {

	// 记录条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 可回收余额信息，如果是余额账户，只会有一条记录。 具体请参见表2。
	AvailRetrieveCoupons *[]AvailRetrieveCoupons `json:"avail_retrieve_coupons,omitempty"`
	HttpStatusCode       int                     `json:"-"`
}

func (o ListMultiAccountRetrieveCouponsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMultiAccountRetrieveCouponsResponse struct{}"
	}

	return strings.Join([]string{"ListMultiAccountRetrieveCouponsResponse", string(data)}, " ")
}
