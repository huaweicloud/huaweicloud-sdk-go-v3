package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPartnerCouponsRecordResponse struct {
	// 查询记录总数。
	TotalCount *int32 `json:"total_count,omitempty"`
	// 日志记录列表。 具体请参见表2。
	Records        *[]CouponRecordV2 `json:"records,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListPartnerCouponsRecordResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPartnerCouponsRecordResponse struct{}"
	}

	return strings.Join([]string{"ListPartnerCouponsRecordResponse", string(data)}, " ")
}
