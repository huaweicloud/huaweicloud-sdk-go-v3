package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportVo 数据导出
type ExportVo struct {

	// **参数解释：** 导出数据类型 **约束限制：** 不涉及 **取值范围：** - export_country_summary：国家及地区统计数据 - export_mgtv_detail：芒果域名统计数据 - export_month_user_charge：月结用户话单计费数据 - export_m5：5分钟粒度统计数据 **默认取值：** 不涉及
	Action *string `json:"action,omitempty"`

	// **参数解释：** 域名列表 > 支持同时输入多个域名  **约束限制：** 不涉及 **取值范围：** - 多个域名用半角逗号（,）分隔 - 如果该参数为all，则为账号下的所有域名报表 **默认取值：** 不涉及
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释：** 导出起始时间 **约束限制：** 不涉及 **取值范围：** 相对于UTC 1970-01-01到当前时间相隔的毫秒数 **默认取值：** 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释：** 导出结束时间 **约束限制：** 不涉及 **取值范围：** 相对于UTC 1970-01-01到当前时间相隔的毫秒数 **默认取值：** 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释：** 数据分组方式 **约束限制：** 不涉及 **取值范围：** - 当action为export_country_summary时，该参数生效 - country：按国家及地区分组 **默认取值：** 不涉及
	GroupBy *string `json:"group_by,omitempty"`

	// **参数解释：** 服务范围 **约束限制：** 服务范围为中国大陆或全球时，加速域名需要到工信部备案 **取值范围：** - mainland_china：中国大陆 - outside_mainland_china：中国大陆境外 - global：全球 **默认取值：** mainland_china：中国大陆
	ServiceArea *string `json:"service_area,omitempty"`

	// **参数解释：** 统计指标类型 **约束限制：** 不涉及 **取值范围：** - flux：流量 - req_num：请求总数 **默认取值：** 不涉及
	StatType *string `json:"stat_type,omitempty"`

	// **参数解释：** 国家&地区编码 **约束限制：** - 查询运营商统计数据时，不传该参数 - 查询top_url数据时，不传该参数 - 查询区域情况数据时，该参数只能传cn（中国）  **取值范围：** - 多个以英文逗号分隔 - all表示全部，取值见附录 **默认取值：** 不涉及
	Country *string `json:"country,omitempty"`
}

func (o ExportVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportVo struct{}"
	}

	return strings.Join([]string{"ExportVo", string(data)}, " ")
}
