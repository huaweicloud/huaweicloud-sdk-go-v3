package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCouponQuotasRecordsResponse Response Object
type ListCouponQuotasRecordsResponse struct {

	// 返回总条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 记录列表。 具体请参见表2。
	Records        *[]QuotaRecord `json:"records,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListCouponQuotasRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCouponQuotasRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListCouponQuotasRecordsResponse", string(data)}, " ")
}
