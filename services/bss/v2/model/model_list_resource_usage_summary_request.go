package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceUsageSummaryRequest Request Object
type ListResourceUsageSummaryRequest struct {

	// 语言。en_US：英文；zh_CN：中文。默认：zh_CN：中文
	XLanguage *string `json:"X-Language,omitempty"`

	// 账期，东八区时间，格式为yyyy-MM。
	BillCycle string `json:"bill_cycle"`

	// 云服务类型，当前仅支持：hws.service.type.cdn：内容分发网络,hws.service.type.obs：对象存储服务,hws.service.type.vpc：虚拟私有云,hws.service.type.iec：智能边缘云
	ServiceTypeCode string `json:"service_type_code"`

	// 资源类型编码，当前仅支持：hws.resource.type.cdn：CDN,hws.resource.type.obs：云存储,hws.resource.type.bandwidth：固定带宽,hws.resource.type.edgebandwidth：边缘固定带宽.资源类型和云服务类型的对应关系可调用[根据云服务类型查询资源列表](https://support.huaweicloud.com/api-oce/qct_00003.html)接口获取。
	ResourceTypeCode string `json:"resource_type_code"`

	// 使用量类型编码，目前仅支持：95Peak：中国大陆月95峰值带宽_1024进制,95peak_1000：中国大陆月95峰值带宽_1000进制,bandwidth95peak：95峰值带宽,95peak_plus：增强型95峰值.资源类型和使用量类型的对应关系可调用[查询使用量类型列表接](https://support.huaweicloud.com/api-oce/qct_00004.html)口获取。
	UsageType string `json:"usage_type"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的数量限制。默认值为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListResourceUsageSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceUsageSummaryRequest struct{}"
	}

	return strings.Join([]string{"ListResourceUsageSummaryRequest", string(data)}, " ")
}
