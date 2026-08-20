package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBandwidthCalcRequest Request Object
type ShowBandwidthCalcRequest struct {

	// **参数解释：** 查询起始时间戳 **约束限制：** 需与结束时间戳同时指定，左闭右开 **取值范围：** - 若查询5分钟时间粒度（即interval为300）数据，start_time设置为整5分钟时刻点，如：1631240100000(对应2021-09-10 10:15:00) - 若查询1小时时间粒度（即interval为3600）数据，start_time设置为整小时时刻点，如：1631239200000(对应2021-09-10 10:00:00) - 若查询1天时间粒度（即interval为86400）数据，start_time设置为东8区零点时刻点，如：1631203200000(对应2021-09-10 00:00:00) **默认取值：** 不涉及
	StartTime int64 `json:"start_time"`

	// **参数解释：** 查询结束时间戳 **约束限制：** 需与起始时间戳同时指定，左闭右开 **取值范围：** - 若查询5分钟时间粒度（即interval为300）数据，end_time设置为整5分钟时刻点，如：1631240100000）对应2021-09-10 10:15:00） - 若查询1小时时间粒度（即interval为3600）数据，end_time设置为整小时时刻点，如：1631239200000（对应2021-09-10 10:00:00） - 若查询1天时间粒度（即interval为86400）数据，end_time设置为东8区零点时刻点，如：1631203200000（对应2021-09-10 00:00:00） **默认取值：** 不涉及
	EndTime int64 `json:"end_time"`

	// **参数解释：** 域名列表 > 如果域名在查询时间段内无数据，结果将不返回该域名的信息  **约束限制：** 仅支持查询已经在CDN创建成功的域名 **取值范围：** - all表示查询名下全部域名 - 多个域名以逗号（半角）分隔，如：www.test1.com,www.test2.com **默认取值：** 不涉及
	DomainName string `json:"domain_name"`

	// **参数解释：** 服务范围 **约束限制：** 当查询回源类指标时，该参数无效 **取值范围：** - mainland_china：中国大陆 - outside_mainland_china：中国大陆境外 **默认取值：** - mainland_china：中国大陆
	ServiceArea *string `json:"service_area,omitempty"`

	// **参数解释：** 企业项目id > 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id  **约束限制：** - 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目 - 当使用子账号调用接口时，该参数必传 **取值范围：** all表示所有项目 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 查询类别，目前支持bw_95（95峰值带宽），bw_peak（日峰值月平均带宽），bw_95_average(95月平均带宽)
	CalcType string `json:"calc_type"`
}

func (o ShowBandwidthCalcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBandwidthCalcRequest struct{}"
	}

	return strings.Join([]string{"ShowBandwidthCalcRequest", string(data)}, " ")
}
