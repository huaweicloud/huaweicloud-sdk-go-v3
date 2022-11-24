package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCouponQuotasReqExt struct {

	// 优惠券额度ID列表。 此参数不携带或携带值为空列表或携带值为null时，不作为筛选条件。
	QuotaIds *[]string `json:"quota_ids,omitempty"`

	// 优惠券额度状态列表。 0：正常3：失效（过期失效和人工设置失效）4：额度调整中（伙伴可以查看该额度，但不能使用该额度发放优惠券）5：冻结 此参数不携带或携带值为空列表或携带值为null时，不作为筛选条件。
	QuotaStatusList *[]int32 `json:"quota_status_list,omitempty"`

	// 优惠券额度的类型。 0：代金券额度1：现金券额度 此参数不携带或携带值为null时，默认值为“0：代金券额度”。
	QuotaType *int32 `json:"quota_type,omitempty"`

	// 创建时间（开始）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出创建时间大于这个时间的记录。 此参数不携带或携带值为null时，不作为筛选条件。
	CreateTimeBegin *string `json:"create_time_begin,omitempty"`

	// 创建时间（结束）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出创建时间小于这个时间的记录。 此参数不携带或携带值为null时，不作为筛选条件。
	CreateTimeEnd *string `json:"create_time_end,omitempty"`

	// 生效时间（开始）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出生效时间大于这个时间的记录。 此参数不携带或携带值为null时，不作为筛选条件。
	EffectiveTimeBegin *string `json:"effective_time_begin,omitempty"`

	// 生效时间（结束）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出生效时间小于这个时间的记录。 此参数不携带或携带值为null时，不作为筛选条件。
	EffectiveTimeEnd *string `json:"effective_time_end,omitempty"`

	// 失效时间（开始）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出失效时间大于这个时间的记录。 此参数不携带或携带值为null时，不作为筛选条件。
	ExpireTimeBegin *string `json:"expire_time_begin,omitempty"`

	// 失效时间（结束）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出失效时间小于这个时间的记录。 此参数不携带或携带值为null时，不作为筛选条件。
	ExpireTimeEnd *string `json:"expire_time_end,omitempty"`

	// 偏移量，从0开始。默认值为0。  说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。 例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询记录数。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 云经销商（二级经销商）ID。 华为云总经销商（一级经销商）查询云经销商的优惠券额度时，需要携带该参数；否则只能查询自己的优惠券额度。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o QueryCouponQuotasReqExt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCouponQuotasReqExt struct{}"
	}

	return strings.Join([]string{"QueryCouponQuotasReqExt", string(data)}, " ")
}
