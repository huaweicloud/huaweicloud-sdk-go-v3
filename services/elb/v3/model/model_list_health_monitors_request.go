package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHealthMonitorsRequest struct {
	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。 使用说明：必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 健康检查ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx****。

	Id *[]string `json:"id,omitempty"`
	// 健康检查端口号。  支持多值查询，查询条件格式：***monitor_port=xxx&monitor_port=xxx***。

	MonitorPort *[]int32 `json:"monitor_port,omitempty"`
	// 发送健康检查请求的域名。  取值：以数字或字母开头，只能包含数字、字母、'-'、'.'。  支持多值查询，查询条件格式：**domain_name=xxx&domain_name=xxx**。

	DomainName *[]string `json:"domain_name,omitempty"`
	// 健康检查名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。

	Name *[]string `json:"name,omitempty"`
	// 健康检查间隔。取值：1-50s。  支持多值查询，查询条件格式：*delay=xxx&delay=xxx*。

	Delay *[]int32 `json:"delay,omitempty"`
	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由OFFLINE判定为ONLINE。取值范围：1-10。  支持多值查询，查询条件格式：*******max_retries=xxx&max_retries=xxx*******。

	MaxRetries *[]int32 `json:"max_retries,omitempty"`
	// 健康检查的管理状态。取值： - true：表示开启健康检查，默认为true。 - false表示关闭健康检查。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 健康检查连续失败多少次后，将后端服务器的健康检查状态由ONLINE判定为OFFLINE。取值范围：1-10。  支持多值查询，查询条件格式：******max_retries_down=xxx&max_retries_down=xxx******。

	MaxRetriesDown *[]int32 `json:"max_retries_down,omitempty"`
	// 一次健康检查请求的超时时间。

	Timeout *int32 `json:"timeout,omitempty"`
	// 健康检查请求协议。 取值：TCP、UDP_CONNECT、HTTP、HTTPS。  支持多值查询，查询条件格式：*****type=xxx&type=xxx*****。

	Type *[]string `json:"type,omitempty"`
	// 期望响应状态码。支持多种取值格式：   单值：单个返回码，例如200。   列表：多个特定返回码，例如200，202。   区间：一个返回码区间，例如200-204。 仅支持HTTP/HTTPS设置该字段，其他协议设置不会生效。 支持多值查询，查询条件格式：****expected_codes=xxx&expected_codes=xxx****。

	ExpectedCodes *[]string `json:"expected_codes,omitempty"`
	// 健康检查测试member健康时发送的http请求路径。默认为\"/\"。  使用说明： - 以\"/\"开头。仅当type为HTTP时生效。  支持多值查询，查询条件格式：***url_path=xxx&url_path=xxx***。

	UrlPath *[]string `json:"url_path,omitempty"`
	// HTTP请求方法，取值：GET、HEAD、POST、PUT、DELETE、TRACE、OPTIONS、CONNECT、PATCH。  支持多值查询，查询条件格式：**http_method=xxx&http_method=xxx**。  不支持该字段，请勿使用。

	HttpMethod *[]string `json:"http_method,omitempty"`
	// 企业项目ID。  支持多值查询，查询条件格式：*enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListHealthMonitorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthMonitorsRequest struct{}"
	}

	return strings.Join([]string{"ListHealthMonitorsRequest", string(data)}, " ")
}
