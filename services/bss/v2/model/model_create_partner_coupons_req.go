package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type CreatePartnerCouponsReq struct {

	// 优惠券额度ID。该值在查询优惠券额度接口的响应参数中获取。
	QuotaId string `json:"quota_id"`

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。
	CustomerIds []string `json:"customer_ids"`

	// 代金券面值。 单位：元。取值大于0且精确到小数点后2位。
	FaceValue *decimal.Decimal `json:"face_value"`

	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 此参数不携带或携带值为null时，赋值为发放优惠券额度的生效时间。
	ValidTime *string `json:"valid_time,omitempty"`

	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 此参数不携带或携带值为null时，赋值为发放优惠券额度的失效时间。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 允许使用的云服务列表，您可以调用查询云服务类型列表接口获取。 多个云服务产品以“,”隔开，最多支持10个。 默认：空（继承额度上的限制）  说明： 如果额度上有限制云服务类型列表，则优惠券上的限制不能超过额度的限制。如果额度上没有限制，则优惠券上可以随意指定云服务类型。
	CloudServiceTypes *[]string `json:"cloud_service_types,omitempty"`

	// 允许使用的产品列表。 多个产品以“,”隔开，最多支持10个。 默认：空（继承额度上的限制）  说明： 如果额度上有限制产品列表，则优惠券上的限制不能超过额度的限制。如果额度上没有限制，则优惠券上可以随意指定产品ID。 产品ID需要合作伙伴通过线下获得。
	ProductIds *[]string `json:"product_ids,omitempty"`

	// 发券时的备注信息。 此参数不携带或携带值为null时，不赋值；携带值为空串时，赋值空串。
	Memo *string `json:"memo,omitempty"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。云经销商给子客户发放优惠券时，需要携带该参数。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o CreatePartnerCouponsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePartnerCouponsReq struct{}"
	}

	return strings.Join([]string{"CreatePartnerCouponsReq", string(data)}, " ")
}
