package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCouponQuotasRecordsRequest struct {
	// 精英服务商ID。获取方法请参见查询精英服务商列表。为空表示查询所有的代金券额度发放回收记录。不为空表示仅查询与该精英服务商相关的代金券额度发放回收记录。默认查询所有精英服务商的代金券额度发放回收记录。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
	// 精英服务商的代金券额度ID。获取方法请参见查询优惠券额度。即华为云伙伴能力中心给精英服务商发放代金券额度时，产生的精英服务商的代金券额度ID，或者从精英服务商回收代金券额度时，精英服务商的代金券额度ID。

	QuotaId *string `json:"quota_id,omitempty"`
	// 查询条件：操作起始时间。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。

	OperationTimeBegin *string `json:"operation_time_begin,omitempty"`
	// 查询条件：操作截止时间。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。

	OperationTimeEnd *string `json:"operation_time_end,omitempty"`
	// 父额度ID。这即华为云伙伴能力中心给精英服务商发放代金券额度时，华为云伙伴能力中心的额度ID，或者从精英服务商回收代金券额度时，回收的华为云伙伴能力中心的额度ID。

	ParentQuotaId *string `json:"parent_quota_id,omitempty"`
	// 操作类型。10：发放额度11：回收额度

	OperationType *string `json:"operation_type,omitempty"`
	// 偏移量，从0开始，默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。

	Offset *int32 `json:"offset,omitempty"`
	// 每次查询的数目。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCouponQuotasRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCouponQuotasRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListCouponQuotasRecordsRequest", string(data)}, " ")
}
