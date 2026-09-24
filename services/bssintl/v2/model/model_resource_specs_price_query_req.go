package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceSpecsPriceQueryReq struct {

	// 云服务类型编码，非必填，范围1-64，此参数不携带或携带值为null时，不作为筛选条件。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 资源类型编码，非必填，范围1-64，此参数不携带或携带值为null时，不作为筛选条件。
	ResourceType *string `json:"resource_type,omitempty"`

	// 区域编码，必填，范围1-64。
	RegionCode string `json:"region_code"`

	// 过滤条件列表，非必填，最多1个。此参数不携带或携带值为空列表或携带值为null时，不作为筛选条件。
	Filters *[]ResourceSpecsPriceFilter `json:"filters,omitempty"`

	// 是否返回资源规格属性信息，非必填，false：不返回（默认）true：返回
	NeedAttributes *bool `json:"need_attributes,omitempty"`

	// 是否返回资源规格官网定价信息，非必填，false：不返回（默认）true：返回
	NeedPrice *bool `json:"need_price,omitempty"`

	// 翻页信息，非必填，首页查询不携带此参数或携带值为null，非首页查询传入上一页响应返回的next_marker
	Marker *string `json:"marker,omitempty"`

	// 查询条数，非必填，取值范围1-50，默认值50
	Limit *int32 `json:"limit,omitempty"`
}

func (o ResourceSpecsPriceQueryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSpecsPriceQueryReq struct{}"
	}

	return strings.Join([]string{"ResourceSpecsPriceQueryReq", string(data)}, " ")
}
