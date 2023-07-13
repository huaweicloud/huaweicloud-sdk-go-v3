package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssuedCouponQuotasRequest Request Object
type ListIssuedCouponQuotasRequest struct {

	// 云经销商的代金券额度ID。获取方法请参见查询优惠券额度。此参数不携带或携带值为空时，不作为筛选条件。
	QuotaId *string `json:"quota_id,omitempty"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。如果需要查询云经销商伙伴的代金券额度，必须携带该字段。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`

	// 父额度ID，即华为云总经销商用于发放给云经销商代金券额度的额度ID。此参数不携带时，不作为筛选条件；携带值为空或携带值为空串时，作为筛选条件。
	ParentQuotaId *string `json:"parent_quota_id,omitempty"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询记录数。默认值为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListIssuedCouponQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssuedCouponQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListIssuedCouponQuotasRequest", string(data)}, " ")
}
