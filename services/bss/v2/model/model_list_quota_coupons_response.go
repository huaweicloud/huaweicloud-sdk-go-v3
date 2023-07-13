package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotaCouponsResponse Response Object
type ListQuotaCouponsResponse struct {

	// 查询总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 额度记录列表。 具体请参见表1。
	Quotas         *[]CouponQuotaV2 `json:"quotas,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListQuotaCouponsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaCouponsResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaCouponsResponse", string(data)}, " ")
}
