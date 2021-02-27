package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListIssuedCouponQuotasResponse struct {
	// 查询的记录总数。
	TotalCount *int32 `json:"total_count,omitempty"`
	// 额度记录列表。 具体请参见表2。
	Quotas         *[]IssuedCouponQuota `json:"quotas,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListIssuedCouponQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListIssuedCouponQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListIssuedCouponQuotasResponse", string(data)}, " ")
}
