package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDomainStatsRequest struct {
	// 查询类型，可选location_summary,location_detail

	Action string `json:"action"`
	// 查询起始时间戳，必须设为5分钟整时刻点

	StartTime int64 `json:"start_time"`
	// 查询结束时间戳，必须设为5分钟整时刻点，与开始时间戳时间差不可以超过一天

	EndTime int64 `json:"end_time"`
	// 查询间隔，对详情类查询有效，如location_detail

	Interval *int64 `json:"interval,omitempty"`
	// 域名列表，多个域名以逗号（半角）分隔，如：www.test1.com,www.test2.com，all表示查询名下全部域名

	DomainName string `json:"domain_name"`
	// 网络资源消耗： - bw（带宽） - flux（流量）  HTTP状态码（组合指标）： - status_code_2xx（状态码详情2xx） - status_code_3xx（状态码详情3xx） - status_code_4xx（状态码详情4xx） - status_code_5xx（状态码详情5xx） - bs_status_code_2xx（回源状态码详情2xx） - bs_status_code_3xx（回源状态码详情3xx） - bs_status_code_4xx（回源状态码详情4xx） - bs_status_code_5xx（回源状态码详情5xx） - status_code和bs_status_code不能一起查询，否则数据会不准确，status_code不支持指定服务区域

	StatType string `json:"stat_type"`
	// 数据分组方式，多个以英文逗号分隔，可选domain,country,district,isp，默认不分组

	GroupBy *string `json:"group_by,omitempty"`
	// 需要过滤的国家编码，多个以英文逗号分隔，all表示全部

	Country *string `json:"country,omitempty"`
	// 需要过滤的地区编码，多个以英文逗号分隔，all表示全部，仅仅country字段为cn时设置才合法

	District *string `json:"district,omitempty"`
	// 需要过滤的运营商编码，多个以英文逗号分隔，all表示全部

	Isp *string `json:"isp,omitempty"`
	// 当用户开启企业项目功能时，该参数生效，表示资源所属企业项目，不传表示默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowDomainStatsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainStatsRequest", string(data)}, " ")
}
