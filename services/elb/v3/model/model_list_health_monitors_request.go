package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHealthMonitorsRequest Request Object
type ListHealthMonitorsRequest struct {

	// **参数解释**：上一页最后一条记录的ID。  **约束限制**： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。  **取值范围**：不涉及  **默认取值**：不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：每页返回的个数。  **约束限制**：不涉及  **取值范围**：0-2000  **默认取值**：2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：是否反向查询。  **约束限制**： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。  **取值范围**： - true：查询上一页。 - false：查询下一页。  **默认取值**：false
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：健康检查ID。 支持多值查询，查询条件格式：*id=xxx&id=xxx****。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Id *[]string `json:"id,omitempty"`

	// **参数解释**：健康检查端口号。 支持多值查询，查询条件格式：***monitor_port=xxx&monitor_port=xxx***。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	MonitorPort *[]int32 `json:"monitor_port,omitempty"`

	// **参数解释**：发送健康检查请求的域名。 支持多值查询，查询条件格式：**domain_name=xxx&domain_name=xxx**。  **约束限制**：不涉及  **取值范围**：以数字或字母开头，只能包含数字、字母、’-’、’.’。  **默认取值**：不涉及
	DomainName *[]string `json:"domain_name,omitempty"`

	// **参数解释**：健康检查名称。 支持多值查询，查询条件格式：*name=xxx&name=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *[]string `json:"name,omitempty"`

	// **参数解释**：健康检查间隔。 支持多值查询，查询条件格式：*delay=xxx&delay=xxx*。  **约束限制**：不涉及  **取值范围**：1-50，单位：秒。  **默认取值**：不涉及
	Delay *[]int32 `json:"delay,omitempty"`

	// **参数解释**：健康检查连续成功多少次后，将后端服务器的健康检查状态由OFFLINE判定为ONLINE。 支持多值查询，查询条件格式：*******max_retries=xxx&max_retries=xxx*******。  **约束限制**：不涉及  **取值范围**：1-10  **默认取值**：不涉及
	MaxRetries *[]int32 `json:"max_retries,omitempty"`

	// **参数解释**：健康检查的管理状态。  **约束限制**：不涉及  **取值范围**： - true：表示开启健康检查。 - false表示关闭健康检查。  **默认取值**：不涉及
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// **参数解释**：健康检查连续失败多少次后，将后端服务器的健康检查状态由ONLINE判定为OFFLINE。 支持多值查询，查询条件格式：******max_retries_down=xxx&max_retries_down=xxx******。  **约束限制**：不涉及  **取值范围**：1-10  **默认取值**：不涉及
	MaxRetriesDown *[]int32 `json:"max_retries_down,omitempty"`

	// **参数解释**：一次健康检查请求的超时时间。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Timeout *int32 `json:"timeout,omitempty"`

	// **参数解释**：健康检查请求协议。 支持多值查询，查询条件格式：*****type=xxx&type=xxx*****。  **约束限制**：不涉及  **取值范围**：TCP、UDP_CONNECT、HTTP、HTTPS、TLS、GRPC[和GRPCS](tag:not_open)。  **默认取值**：不涉及
	Type *[]string `json:"type,omitempty"`

	// **参数解释**：期望响应状态码。 支持多值查询，查询条件格式：****expected_codes=xxx&expected_codes=xxx****。  **约束限制**：不涉及  **取值范围**： - 单值：单个返回码，例如200。 - 列表：多个特定返回码，例如200，202。 - 区间：一个返回码区间，例如200-204。  **默认取值**：不涉及
	ExpectedCodes *[]string `json:"expected_codes,omitempty"`

	// **参数解释**：健康检查测试member健康时发送的http请求路径。 支持多值查询，查询条件格式：***url_path=xxx&url_path=xxx***。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	UrlPath *[]string `json:"url_path,omitempty"`

	// **参数解释**：HTTP请求方法。 支持多值查询，查询条件格式：**http_method=xxx&http_method=xxx**。  **约束限制**：不涉及  **取值范围**：GET、HEAD、POST  **默认取值**：不涉及
	HttpMethod *[]string `json:"http_method,omitempty"`

	// **参数解释**：资源所属的企业项目ID。 支持多值查询，查询条件格式：*enterprise_project_id=xxx&enterprise_project_id=xxx*。  **约束限制**： - 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:healthmonitors:list权限。 - 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  **取值范围**：不涉及  **默认取值**：不涉及  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListHealthMonitorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthMonitorsRequest struct{}"
	}

	return strings.Join([]string{"ListHealthMonitorsRequest", string(data)}, " ")
}
