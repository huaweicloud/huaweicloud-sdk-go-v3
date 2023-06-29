package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIncentiveDiscountPoliciesRequest Request Object
type ListIncentiveDiscountPoliciesRequest struct {

	// 查询策略的指定时间。东八区时间，格式：YYYY-MM。 说明： 实际查询结果为指定时间所在月最后一天23:59:59的策略情况。
	Time string `json:"time"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串时，作为筛选条件。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的数量，默认值为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListIncentiveDiscountPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIncentiveDiscountPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListIncentiveDiscountPoliciesRequest", string(data)}, " ")
}
