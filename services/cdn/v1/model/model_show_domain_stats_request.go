package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDomainStatsRequest struct {
	// 查询类型，可选location_summary,location_detail  location_summary：查询汇总数据 location_detail：查询数据详情

	Action string `json:"action"`
	// 查询起始时间戳， 时间戳应设置需为整5分钟或整小时时刻点，设置方式如下  interval为300时，start_time设置为整5分钟时刻点，如：1631240100000(对应2021-09-10 10:15:00) interval大于等于3600时，start_time设置为整小时时刻点，如：1631239200000(对应2021-09-10 10:00:00)

	StartTime int64 `json:"start_time"`
	// 查询结束时间戳， 时间戳应设置需为整5分钟或整小时时刻点，设置方式如下  interval为300时，end_time设置为整5分钟时刻点，如：1631243700000(对应2021-09-11 10:15:00) interval大于等于3600时，end_time设置为整小时时刻点，如：1631325600000(对应2021-09-11 10:00:00)

	EndTime int64 `json:"end_time"`
	// 查询时间间隔，单位为秒，可设置值300(5分钟),3600(1小时),14400(4小时)等

	Interval *int64 `json:"interval,omitempty"`
	// 域名列表，多个域名以逗号（半角）分隔，如：www.test1.com,www.test2.com，all表示查询名下全部域名

	DomainName string `json:"domain_name"`
	// 网络资源消耗： - bw（带宽） - flux（流量）  HTTP状态码（组合指标）： - status_code_2xx（状态码详情2xx） - status_code_3xx（状态码详情3xx） - status_code_4xx（状态码详情4xx） - status_code_5xx（状态码详情5xx）

	StatType string `json:"stat_type"`
	// 数据分组方式，多个以英文逗号分隔，可选domain,country,district,isp，默认不分组

	GroupBy *string `json:"group_by,omitempty"`
	// 需要过滤的国家编码，多个以英文逗号分隔，all表示全部

	Country *string `json:"country,omitempty"`
	// 需要过滤的地区编码，多个以英文逗号分隔，all表示全部，仅仅country字段为cn时设置才合法

	District *string `json:"district,omitempty"`
	// 需要过滤的运营商编码，多个以英文逗号分隔，all表示全部

	Isp *string `json:"isp,omitempty"`
	// 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目，不传表示查询默认项目。注意：当使用子账号调用接口时，该参数必传。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowDomainStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainStatsRequest", string(data)}, " ")
}
