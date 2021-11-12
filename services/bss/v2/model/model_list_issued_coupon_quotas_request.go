package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIssuedCouponQuotasRequest struct {
	// 精英服务商的代金券额度ID。

	QuotaId *string `json:"quota_id,omitempty"`
	// 精英服务商ID。获取方法请参见查询精英服务商列表。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
	// 父额度ID，即华为云伙伴能力中心用于发放给精英服务商代金券额度的额度ID。

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
