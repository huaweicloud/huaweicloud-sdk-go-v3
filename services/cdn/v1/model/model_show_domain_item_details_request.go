package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDomainItemDetailsRequest struct {
	// 当用户开启企业项目功能时，该参数生效，表示资源所属企业项目，不传表示默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 查询起始时间戳，必须设为5分钟整时刻点

	StartTime int64 `json:"start_time"`
	// 查询结束时间戳，必须设为5分钟整时刻点，与开始时间戳时间差不可以超过一天

	EndTime int64 `json:"end_time"`
	// 域名列表，多个域名以逗号（半角）分隔，如：www.test1.com,www.test2.com，all表示查询名下全部域名

	DomainName string `json:"domain_name"`
	// mainland_china(中国大陆)，outside_mainland_china(中国大陆境外)，默认为mainland_china。

	ServiceArea *string `json:"service_area,omitempty"`
	// 网络资源消耗： - bw（带宽） - flux（流量） - bs_bw(回源带宽) - bs_flux（回源流量）  访问情况： - req_num（请求总数） - hit_num（请求命中次数） - bs_num(回源总数) - bs_fail_num(回源失败数) - hit_flux（命中流量）  HTTP状态码（组合指标）： - http_code_2xx(状态码汇总2xx) - http_code_3xx(状态码汇总3xx) - http_code_4xx(状态码汇总4xx) - http_code_5xx(状态码汇总5xx) - bs_http_code_2xx（回源状态码汇总2xx） - bs_http_code_3xx（回源状态码汇总3xx） - bs_http_code_4xx（回源状态码汇总4xx） - bs_http_code_5xx（回源状态码汇总5xx） - status_code_2xx（状态码详情2xx） - status_code_3xx（状态码详情3xx） - status_code_4xx（状态码详情4xx） - status_code_5xx（状态码详情5xx） - bs_status_code_2xx（回源状态码详情2xx） - bs_status_code_3xx（回源状态码详情3xx） - bs_status_code_4xx（回源状态码详情4xx） - bs_status_code_5xx（回源状态码详情5xx） - status_code和bs_status_code不能一起查询，否则数据会不准确，status_code不支持指定服务区域

	StatType string `json:"stat_type"`
}

func (o ShowDomainItemDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainItemDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainItemDetailsRequest", string(data)}, " ")
}
