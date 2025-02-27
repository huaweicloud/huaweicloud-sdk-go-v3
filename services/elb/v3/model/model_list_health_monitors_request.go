package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHealthMonitorsRequest Request Object
type ListHealthMonitorsRequest struct {

	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。
	Marker *string `json:"marker,omitempty"`

	// 参数解释：每页返回的个数。  取值范围：0-2000  默认取值：2000
	Limit *int32 `json:"limit,omitempty"`

	// 是否反向查询。  取值： - true：查询上一页。 - false：查询下一页，默认。  使用说明： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 健康检查ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx****。
	Id *[]string `json:"id,omitempty"`

	// 健康检查端口号。  支持多值查询，查询条件格式：***monitor_port=xxx&monitor_port=xxx***。
	MonitorPort *[]int32 `json:"monitor_port,omitempty"`

	// 发送健康检查请求的域名。  取值：以数字或字母开头，只能包含数字、字母、’-’、’.’。  支持多值查询，查询条件格式：**domain_name=xxx&domain_name=xxx**。
	DomainName *[]string `json:"domain_name,omitempty"`

	// 健康检查名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。
	Name *[]string `json:"name,omitempty"`

	// 健康检查间隔。  取值：1-50s。  支持多值查询，查询条件格式：*delay=xxx&delay=xxx*。
	Delay *[]int32 `json:"delay,omitempty"`

	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由OFFLINE判定为ONLINE。取值范围：1-10。  支持多值查询，查询条件格式：*******max_retries=xxx&max_retries=xxx*******。
	MaxRetries *[]int32 `json:"max_retries,omitempty"`

	// 参数解释：健康检查的管理状态。  取值范围： - true：表示开启健康检查。 - false表示关闭健康检查。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 健康检查连续失败多少次后，将后端服务器的健康检查状态由ONLINE判定为OFFLINE。取值范围：1-10。  支持多值查询，查询条件格式：******max_retries_down=xxx&max_retries_down=xxx******。
	MaxRetriesDown *[]int32 `json:"max_retries_down,omitempty"`

	// 一次健康检查请求的超时时间。
	Timeout *int32 `json:"timeout,omitempty"`

	// 健康检查请求协议。  取值：TCP、UDP_CONNECT、HTTP、HTTPS、TLS和GRPC。  支持多值查询，查询条件格式：*****type=xxx&type=xxx*****。
	Type *[]string `json:"type,omitempty"`

	// 期望响应状态码。  取值： - 单值：单个返回码，例如200。 - 列表：多个特定返回码，例如200，202。 - 区间：一个返回码区间，例如200-204。   默认值：200。  仅支持HTTP/HTTPS/GRPC设置该字段，其他协议设置不会生效。  支持多值查询，查询条件格式：****expected_codes=xxx&expected_codes=xxx****。
	ExpectedCodes *[]string `json:"expected_codes,omitempty"`

	// 健康检查测试member健康时发送的http请求路径。默认为“/”。  使用说明：以“/”开头。当type为HTTP/HTTPS时生效。  支持多值查询，查询条件格式：***url_path=xxx&url_path=xxx***。
	UrlPath *[]string `json:"url_path,omitempty"`

	// HTTP请求方法。  取值：GET、HEAD、POST。  支持多值查询，查询条件格式：**http_method=xxx&http_method=xxx**。
	HttpMethod *[]string `json:"http_method,omitempty"`

	// 参数解释：所属的企业项目ID。 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:healthmonitors:list权限。 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  支持多值查询，查询条件格式： *enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListHealthMonitorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthMonitorsRequest struct{}"
	}

	return strings.Join([]string{"ListHealthMonitorsRequest", string(data)}, " ")
}
