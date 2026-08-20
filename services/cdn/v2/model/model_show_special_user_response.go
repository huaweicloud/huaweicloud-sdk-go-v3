package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpecialUserResponse Response Object
type ShowSpecialUserResponse struct {

	// **参数解释：** 租户是否可以查询总请求时长 **约束限制：** 不涉及 **取值范围：** - 1：表示租户可以查询总请求时长 - 0：表示租户不可以查询总请求时长 **默认取值：** 不涉及
	Status *int64 `json:"status,omitempty"`

	// **参数解释：** 进制 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Metric *int64 `json:"metric,omitempty"`

	// **参数解释：** 流量进制 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	FluxMetric *int64 `json:"flux_metric,omitempty"`

	// **参数解释：** 租户是否是开放国家及地区界面 **约束限制：** 不涉及 **取值范围：** - 1：表示租户开放国家及地区界面 - 0：表示租户不用开放国家及地区界面 **默认取值：** 不涉及
	Cy *int64 `json:"cy,omitempty"`

	// **参数解释：** 租户是否可以查询ipv6流量、https流量 **约束限制：** 不涉及 **取值范围：** - 1：表示租户可以查询ipv6流量、https流量 - 0：表示租户不可以查询ipv6流量、https流量 **默认取值：** 不涉及
	H6 *int64 `json:"h6,omitempty"`

	// **参数解释：** 租户是否可以查询具体的状态码详情 **约束限制：** 不涉及 **取值范围：** - 1：表示租户可以查询具体的状态码详情 - 0：表示租户不可以查询具体的状态码详情 **默认取值：** 不涉及
	C *int64 `json:"c,omitempty"`

	// **参数解释：** 查询查询top url时是否返回http状态码 **约束限制：** 不涉及 **取值范围：** - 1：表示租户查询top url时返回http状态码 - 0：表示租户查询top url时不返回http状态码 **默认取值：** 不涉及
	Sc *int64 `json:"sc,omitempty"`

	// **参数解释：** 租户是否可以查询回源状态码 **约束限制：** 不涉及 **取值范围：** - 1：表示租户可以查询回源状态码 - 0：表示租户不可以查询回源状态码 **默认取值：** 不涉及
	Bhc *int64 `json:"bhc,omitempty"`

	// **参数解释：** 租户是否可以查询protocol和IPVersion **约束限制：** 不涉及 **取值范围：** - 1：表示租户可以查询protocol和IPVersion - 0：表示租户不可以查询protocol和IPVersion **默认取值：** 不涉及
	Pi *int64 `json:"pi,omitempty"`

	// **参数解释：** 租户是否可以导出5分钟粒度数据 **约束限制：** 不涉及 **取值范围：** - 1：表示租户可以导出5分钟粒度数据 - 0：表示租户不可以导出5分钟粒度数据 **默认取值：** 不涉及
	Exp5 *int64 `json:"exp5,omitempty"`

	// **参数解释：** 租户是否可以查询1分钟粒度数据 **约束限制：** 不涉及 **取值范围：** - 1：表示租户可以查询1分钟粒度数据 - 0：表示租户不可以查询1分钟粒度数据 **默认取值：** 不涉及
	M1 *int64 `json:"m1,omitempty"`

	// **参数解释：** 租户是否可以查询1个月5分钟粒度统计数据 **约束限制：** 不涉及 **取值范围：** - 1：表示租户可以查询1个月5分钟粒度统计数据 - 0：表示租户不可以查询1个月5分钟粒度统计数据 **默认取值：** 不涉及
	IsMonthM5 *int64 `json:"is_month_m5,omitempty"`

	// **参数解释：** 租户是否可以在租户界面指定下载链接可用范围 **约束限制：** 不涉及 **取值范围：** - 1：表示租户可以在租户界面指定下载链接可用范围 - 0：表示租户不可以在租户界面指定下载链接可用范围 **默认取值：** 不涉及
	ExpAgy *int64 `json:"exp_agy,omitempty"`

	// **参数解释：** 租户是否可以上报到国际站CES **约束限制：** 不涉及 **取值范围：** - 1：表示租户可以上报到国际站CES - 0：表示租户不可以上报到国际站CES **默认取值：** 不涉及
	CesReportSite *int64 `json:"ces_report_site,omitempty"`

	// **参数解释：** 租户是否按上浮系数展示数据 **约束限制：** 不涉及 **取值范围：** - 1：表示租户按上浮系数展示数据 - 0：表示租户不按上浮系数展示数据 **默认取值：** 不涉及
	Float *int64 `json:"float,omitempty"`

	// **参数解释：** 租户是否允许查询入网带宽 **约束限制：** 不涉及 **取值范围：** - 1：表示租户允许查询入网带宽 - 0：表示租户不可以查询入网带宽 **默认取值：** 不涉及
	IsShowUpBw     *int64 `json:"is_show_up_bw,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSpecialUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpecialUserResponse struct{}"
	}

	return strings.Join([]string{"ShowSpecialUserResponse", string(data)}, " ")
}
