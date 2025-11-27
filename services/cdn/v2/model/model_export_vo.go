package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportVo 数据导出
type ExportVo struct {

	// **参数解释：** 规则行为 **约束限制：** 不涉及
	Action *string `json:"action,omitempty"`

	// 域名列表，支持同时输入多个域名，多个域名用半角逗号（,）分隔；说明：如果该参数为all，则为账号下的所有域名报表。
	DomainName *string `json:"domain_name,omitempty"`

	// 查询起始时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询结束时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。
	EndTime *int64 `json:"end_time,omitempty"`

	// 数据分组方式
	GroupBy *string `json:"group_by,omitempty"`

	// **参数解释：** 域名服务范围 **约束限制：** 服务范围为中国大陆或全球时，加速域名需要到工信部备案 **取值范围：** - mainland_china: 中国大陆 - outside_mainland_china: 中国大陆境外 - global: 全球 **默认取值：** mainland_china: 中国大陆
	ServiceArea *string `json:"service_area,omitempty"`

	// 参数类型支持：flux(流量)，req_num(请求总数)。
	StatType *string `json:"stat_type,omitempty"`

	// - 国家&地区编码，多个以英文逗号分隔，all表示全部，取值见附录 - 访问运营商统计数据时不能填写 - 访问top_url数据时不能填写 - 访问区域情况数据时只能填写cn(中国)
	Country *string `json:"country,omitempty"`
}

func (o ExportVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportVo struct{}"
	}

	return strings.Join([]string{"ExportVo", string(data)}, " ")
}
