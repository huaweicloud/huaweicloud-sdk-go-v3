package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCustomerOnDemandResourcesReq struct {

	// 客户账号ID。 您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。
	CustomerId string `json:"customer_id"`

	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	RegionCode *string `json:"region_code,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 资源ID批量查询。 用于查询指定资源ID对应的资源。 最多支持同时传递50个ID的列表。 此参数不携带或携带值为空列表或携带值为null时，不作为筛选条件。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 生效时间的开始时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	EffectiveTimeBegin *string `json:"effective_time_begin,omitempty"`

	// 生效时间的结束时间 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	EffectiveTimeEnd *string `json:"effective_time_end,omitempty"`

	// 偏移量，从0开始。默认值为0。  说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。 例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 一次查询的条数，默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 资源状态： 1：正常（已开通）2：宽限期3：冻结中4：变更中5：正在关闭6：已关闭 此参数不携带或携带值为null时，不作为筛选条件。
	Status *int32 `json:"status,omitempty"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。如果需要查询云经销商子客户的按需资源列表，必须携带该字段，除此之外，此参数不做处理。否则只能查询自己的子客户按需资源。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o QueryCustomerOnDemandResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCustomerOnDemandResourcesReq struct{}"
	}

	return strings.Join([]string{"QueryCustomerOnDemandResourcesReq", string(data)}, " ")
}
