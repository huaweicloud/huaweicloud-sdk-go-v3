package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHealthmonitorsRequest Request Object
type ListHealthmonitorsRequest struct {

	// 分页查询中每页的健康检查个数
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询的起始的资源id，表示上一页最后一条查询记录的健康检查的id。不指定时表示查询第一页。
	Marker *string `json:"marker,omitempty"`

	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 健康检查ID。
	Id *string `json:"id,omitempty"`

	// 健康检查名称。
	Name *string `json:"name,omitempty"`

	// 健康检查间隔，单位秒，取值范围[1，50]。
	Delay *int32 `json:"delay,omitempty"`

	// 健康检查最大重试次数，取值范围[1，10]。
	MaxRetries *int32 `json:"max_retries,omitempty"`

	// 健康检查的管理状态。取值范围：true/false。默认为true；true表示开启健康检查；false表示关闭健康检查。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 健康检查超时时间，单位秒，取值范围[1，50]。 建议该值小于delay的值。
	Timeout *int32 `json:"timeout,omitempty"`

	// 健康检查的类型。取值范围：TCP、UDP_CONNECT、HTTP。
	Type *string `json:"type,omitempty"`

	// 健康检查端口号]。默认为空，表示使用后端云服务器的protocol_port作为健康检查的检查端口。
	MonitorPort *int32 `json:"monitor_port,omitempty"`

	// 期望HTTP响应状态码；默认值：“200”。取值范围：单值，例如200；列表，例如200，202；区间，例如200-204。仅当type为HTTP时生效。 该字段为预留字段，暂未启用。
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// 健康检查时，发送的http请求的域名。仅当type为HTTP时生效。默认为空，表示使用负载均衡器的vip_address作为http请求的目的地址。以数字或字母开头，只能包含数字、字母、’-’、’.’。例如：www.huaweitest.com
	DomainName *string `json:"domain_name,omitempty"`

	// 健康检查时发送的http请求路径。默认为“/”。以“/”开头。仅当type为HTTP时生效。例如：“/test”
	UrlPath *string `json:"url_path,omitempty"`

	// HTTP请求的方法；默认值：GET取值范围：GET、HEAD、POST、PUT、DELETE、TRACE、OPTIONS、CONNECT、PATCH。仅当type为HTTP时生效。
	HttpMethod *string `json:"http_method,omitempty"`
}

func (o ListHealthmonitorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthmonitorsRequest struct{}"
	}

	return strings.Join([]string{"ListHealthmonitorsRequest", string(data)}, " ")
}
