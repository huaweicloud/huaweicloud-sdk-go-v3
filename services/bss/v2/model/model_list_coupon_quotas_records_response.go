package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCouponQuotasRecordsResponse struct {

	// 返回总条数。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 记录列表。 具体请参见表2。
	Records        *[]QuotaRecord `json:"records,omitempty" xml:"records"`
	HttpStatusCode int            `json:"-"`
}

func (o ListCouponQuotasRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCouponQuotasRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListCouponQuotasRecordsResponse", string(data)}, " ")
}
