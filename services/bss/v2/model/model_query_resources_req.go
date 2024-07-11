package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryResourcesReq struct {

	// 资源ID列表。 查询指定资源ID的资源（当only_main_resource=0时，查询指定资源及其附属资源）。最大支持50个ID同时作为条件查询。 此参数不携带或携带值为空列表时，不作为筛选条件，返回其他条件匹配的记录。不支持携带值为null。  说明： 资源ID是指开通资源以后，云服务针对该资源分配的标志，譬如云主机ECS的资源ID是server_id。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 订单号。查询指定订单下的资源。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件，返回其他条件匹配的记录。  说明： 使用特殊字符进行查询的时候，请注意进行URL编码转换，如“%”的转码应为“%25”。
	OrderId *string `json:"order_id,omitempty"`

	// 是否只查询主资源，该参数对于请求参数是子资源ID的时候无效，如果resource_ids是子资源ID，只能查询自己。 0：查询主资源及附属资源。1：只查询主资源。 默认值为0。  说明： 主资源是指有关联的几个资源中，处于主导位置的资源。 对于ECS而言，虚拟机VM是主资源，磁盘EVS是辅资源。对于VPC而言，共享带宽的情况下，带宽为主资源，对应的从资源为弹性IP（可能包含多个IP）；独享带宽的情况下，弹性IP为主资源，对应的从资源为带宽。
	OnlyMainResource *int32 `json:"only_main_resource,omitempty"`

	// 资源状态。 查询指定状态的资源。多个状态以英文逗号分隔。 2：使用中3：已关闭4：已冻结5：已过期 此参数不携带或携带值为空列表时，不作为筛选条件，返回所有状态的资源列表。
	StatusList *[]int32 `json:"status_list,omitempty"`

	// 偏移量，从0开始。默认值为0。  说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。 例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的条数。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 查询指定时间段内失效的资源列表，时间段的起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件，返回其他条件匹配的记录。
	ExpireTimeBegin *string `json:"expire_time_begin,omitempty"`

	// 查询指定时间段内失效的资源列表，时间段的结束时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件，返回其他条件匹配的记录。
	ExpireTimeEnd *string `json:"expire_time_end,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。 此参数不携带、携带值为null，不作为筛选条件。此参数不允许为空串，有参数校验。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`
}

func (o QueryResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResourcesReq struct{}"
	}

	return strings.Join([]string{"QueryResourcesReq", string(data)}, " ")
}
