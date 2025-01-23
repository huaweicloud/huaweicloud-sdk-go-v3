package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPartnerCouponsRecordRequest Request Object
type ListPartnerCouponsRecordRequest struct {

	// 操作类型。1：发放2：手动回收3：解绑自动回收4：过期回收5：退订回收6：退款充值撤销7：建立关联回收此参数不携带时，不作为筛选条件；携带值为空列表或携带值为null时，作为筛选条件。
	OperationTypes *[]string `json:"operation_types,omitempty"`

	// 额度ID。请从“查询优惠券额度”接口的响应参数中获取。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串或携带值为null时，作为筛选条件。
	QuotaId *string `json:"quota_id,omitempty"`

	// 额度类型。0：代金券额度1：现金券额度此参数不携带或携带值为空或携带值为null时，不作为筛选条件；不支持携带值为空串。
	QuotaType *int32 `json:"quota_type,omitempty"`

	// 代金券ID列表。请从“发放优惠券”接口的响应参数中获取。此参数不携带时，不作为筛选条件；携带值为空列表或携带值为null时，作为筛选条件。
	CouponIds *[]string `json:"coupon_ids,omitempty"`

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串或携带值为null时，作为筛选条件。
	CustomerId *string `json:"customer_id,omitempty"`

	// 操作时间（开始）。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。输入这个条件，会查询出操作时间大于这个时间的记录。此参数不携带或携带值为空时，不作为筛选条件；不支持携带值为空串和携带值为null。
	OperationTimeBegin *string `json:"operation_time_begin,omitempty"`

	// 操作时间（结束）。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。输入这个条件，会查询出操作时间小于这个时间的记录。此参数不携带或携带值为空时，不作为筛选条件；不支持携带值为空串和携带值为null。
	OperationTimeEnd *string `json:"operation_time_end,omitempty"`

	// 操作结果。0：成功-1：失败（非0的记录）此参数不携带或携带值为非枚举值时，不作为筛选条件。
	Result *string `json:"result,omitempty"`

	// 偏移量，从0开始。默认值为0。此参数不携带或携带值为空或携带值为null时，取默认值0；不支持携带值为空串。说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数。默认值为10。此参数不携带或携带值为空或携带值为null时，取默认值10；不支持携带值为空串。
	Limit *int32 `json:"limit,omitempty"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。华为云总经销商（一级经销商）查询云经销商发放给子客户的优惠券发放记录时，需要携带该参数，否则只能查询发给自己子客户的优惠券发放记录。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListPartnerCouponsRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartnerCouponsRecordRequest struct{}"
	}

	return strings.Join([]string{"ListPartnerCouponsRecordRequest", string(data)}, " ")
}
