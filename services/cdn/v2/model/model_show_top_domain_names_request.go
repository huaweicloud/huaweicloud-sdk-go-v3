package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopDomainNamesRequest Request Object
type ShowTopDomainNamesRequest struct {

	// - 查询起始时间戳，时间戳应设置需为整点时间戳，设置方式如下： - interval为3600时，start_time设置为整小时时刻点，如：1631239200000(对应2021-09-10 10:00:00) - interval为86400时，start_time设置为东8区零点时刻点，如：1631203200000(对应2021-09-10 00:00:00)
	StartTime int64 `json:"start_time"`

	// - 查询结束时间戳，时间戳应设置需为整点时间戳，设置方式如下： - interval为3600时，end_time设置为整小时时刻点，如：1631325600000(对应2021-09-11 10:00:00) - interval为86400时，end_time设置为东8区零点时刻点，如：1631376000000(对应2021-09-12 00:00:00)
	EndTime int64 `json:"end_time"`

	// - 统计类型 - 目前只支持bw（带宽），flux（流量），req_num（请求总数）
	StatType string `json:"stat_type"`

	// **参数解释：** 服务范围 **约束限制：** 当查询回源类指标时，该参数无效 **取值范围：** - mainland_china：中国大陆 - outside_mainland_china：中国大陆境外 **默认取值：** - mainland_china：中国大陆
	ServiceArea *string `json:"service_area,omitempty"`

	// **参数解释：** top域名查询数量 **约束限制：** 不涉及 **取值范围：** 0-500 **默认取值：** 20
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 企业项目id > 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id  **约束限制：** - 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目 - 当使用子账号调用接口时，该参数必传 **取值范围：** all表示所有项目 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowTopDomainNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopDomainNamesRequest struct{}"
	}

	return strings.Join([]string{"ShowTopDomainNamesRequest", string(data)}, " ")
}
