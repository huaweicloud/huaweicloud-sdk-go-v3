package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpecialUserResponse Response Object
type ShowSpecialUserResponse struct {

	// 1表示用户可以查询总请求时长枚举，0表示用户不可以查询总请求时长枚举
	Status *int64 `json:"status,omitempty"`

	// 进制
	Metric *int64 `json:"metric,omitempty"`

	// 流量进制
	FluxMetric *int64 `json:"flux_metric,omitempty"`

	// 1表示用户可以，0表示用户不可以。是否是开放国家及地区界面用户
	Cy *int64 `json:"cy,omitempty"`

	// 1表示用户可以查询ipv6流量,https流量，0表示用户不可以
	H6 *int64 `json:"h6,omitempty"`

	// 1表示用户可以查询具体的状态码详情，0表示用户不可以
	C *int64 `json:"c,omitempty"`

	// 1表示用户查询 top url 时要返回http状态码，0表示用户不返回
	Sc *int64 `json:"sc,omitempty"`

	// 1表示该用户可以查询回源状态码，0表示不可以
	Bhc *int64 `json:"bhc,omitempty"`

	// 1表示该用户可以查询protocol和IPVersion，0表示用户不可以
	Pi *int64 `json:"pi,omitempty"`

	// 1表示该用户可以查询租户界面5分钟粒度数据导出白名单，0表示用户不可以
	Exp5 *int64 `json:"exp5,omitempty"`

	// 1表示该用户可以查询1分钟粒度统计数据，0表示用户不可以
	M1 *int64 `json:"m1,omitempty"`

	// 1表示该用户可以查询1个月5分钟粒度统计数据，0表示用户不可以
	IsMonthM5 *int64 `json:"is_month_m5,omitempty"`

	// 1表示该用户可以在租户界面指定下载链接可用范围，0表示用户不可以
	ExpAgy *int64 `json:"exp_agy,omitempty"`

	// 1表示该用户可以是否上报到国际站CES，0表示用户不可以
	CesReportSite *int64 `json:"ces_report_site,omitempty"`

	// 1表示该用户按上浮系数展示数据，0表示用户不可以
	Float *int64 `json:"float,omitempty"`

	// 1表示该用户允许查询入网带宽，0表示用户不可以
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
