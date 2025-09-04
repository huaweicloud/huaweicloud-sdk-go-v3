package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainClientStatsRequest Request Object
type ListDomainClientStatsRequest struct {

	// 查询起始时间戳，需与结束时间戳同时指定，左闭右开，设置方式如下： - interval为300时，start_time设置为整5分钟时刻点，如：1631240100000(对应2021-09-10 10:15:00) - interval为3600时，start_time设置为整小时时刻点，如：1631239200000(对应2021-09-10 10:00:00) - interval为86400时，start_time设置为东8区零点时刻点，如：1631203200000(对应2021-09-10 00:00:00)
	StartTime int64 `json:"start_time"`

	// 查询结束时间戳，需与开始时间戳同时指定，左闭右开，设置方式如下： - interval为300时，end_time设置为整5分钟时刻点，如：1631243700000(对应2021-09-10 11:15:00) - interval为3600时，end_time设置为整小时时刻点，如：1631325600000(对应2021-09-11 10:00:00) - interval为86400时，end_time设置为东8区零点时刻点，如：1631376000000(对应2021-09-12 00:00:00)
	EndTime int64 `json:"end_time"`

	// 域名列表，多个域名以逗号（半角）分隔，如：www.test1.com,www.test2.com all表示查询名下全部域名。如果域名在查询时间段内无数据，结果将不返回该域名的信息。
	DomainName string `json:"domain_name"`

	// -  客户端访问资源指标类别：  - ip_num (去重后ip访问数量) - uv(五分钟内的访问独立IP数，一个独立ip记一次。)
	StatType string `json:"stat_type"`

	// 服务区域：mainland_china(大陆)，outside_mainland_china(海外)，默认为mainland_china
	ServiceArea *string `json:"service_area,omitempty"`

	// 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目，\"all\"表示所有项目。注意：当使用子账号调用接口时，该参数必传。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListDomainClientStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainClientStatsRequest struct{}"
	}

	return strings.Join([]string{"ListDomainClientStatsRequest", string(data)}, " ")
}
