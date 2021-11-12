package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHealthMonitorsRequest struct {
	// 健康检查的管理状态；该字段虽然支持创建、更新，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。该字段虽然支持创建、更新，但实际取值决定于member对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 健康检查间隔。

	Delay *[]int32 `json:"delay,omitempty"`
	// 功能说明：健康检查测试member健康状态时，发送的http请求的域名。仅当type为HTTP时生效。使用说明：默认为空，表示使用负载均衡器的vip作为http请求的目的地址。以数字或字母开头，只能包含数字、字母、’-’、’.’。

	DomainName *[]string `json:"domain_name,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// 期望HTTP响应状态码，指定下列值：单值，例如200；列表，例如200，202；区间，例如200-204。仅当type为HTTP时生效。该字段为预留字段，暂未启用。

	ExpectedCodes *[]string `json:"expected_codes,omitempty"`
	// HTTP方法，可以为GET、HEAD、POST、PUT、DELETE、TRACE、OPTIONS、CONNECT、PATCH。仅当type为HTTP时生效。该字段为预留字段，暂未启用。

	HttpMethod *[]string `json:"http_method,omitempty"`
	// 健康检查ID。

	Id *[]string `json:"id,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 上一页最后一条记录的ID。  使用说明：  - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 最大重试次数。

	MaxRetries *[]int32 `json:"max_retries,omitempty"`
	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由ONLINE判定为OFFLINE，取值范围[1，10]。

	MaxRetriesDown *[]int32 `json:"max_retries_down,omitempty"`
	// 健康检查端口号。

	MonitorPort *[]int32 `json:"monitor_port,omitempty"`
	// 健康检查名称。

	Name *[]string `json:"name,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。 使用说明：必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 健康检查的超时时间。建议该值小于delay的值。

	Timeout *int32 `json:"timeout,omitempty"`
	// 健康检查类型。

	Type *[]string `json:"type,omitempty"`
	// 功能说明：健康检查测试member健康时发送的http请求路径。默认为“/”。使用说明：以“/”开头。仅当type为HTTP时生效。

	UrlPath *[]string `json:"url_path,omitempty"`
}

func (o ListHealthMonitorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthMonitorsRequest struct{}"
	}

	return strings.Join([]string{"ListHealthMonitorsRequest", string(data)}, " ")
}
