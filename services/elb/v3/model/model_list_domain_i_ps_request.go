package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainIPsRequest Request Object
type ListDomainIPsRequest struct {

	// **参数解释**：负载均衡器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId string `json:"loadbalancer_id"`

	// **参数解释**：上一页最后一条记录的ID。  **约束限制**： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。  **取值范围**：不涉及  **默认取值**：不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：每页返回的个数。  **约束限制**：不涉及  **取值范围**：0-2000  **默认取值**：2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：是否反向查询。  **约束限制**： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。  **取值范围**： - true：查询上一页。 - false：查询下一页。  **默认取值**：false
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：IPv4或IPv6地址。 支持多值查询，查询条件格式： *ip_address=xxx&ip_address=xxx*。  **约束限制**：必须是当前负载均衡器绑定的私网地址或者公网地址。  **取值范围**：不涉及  **默认取值**：不涉及
	IpAddress *[]string `json:"ip_address,omitempty"`

	// **参数解释**：IP地址是否已加入到域名解析。  **约束限制**：不涉及  **取值范围**： - true：已加入域名解析。 - false：未加入域名解析。  **默认取值**：不涉及
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**：IP地址类型。 支持多值查询，查询条件格式： *type=xxx&type=xxx*。  **约束限制**：不涉及  **取值范围**： - vip：私网IP。 - eip：公网IP。  **默认取值**：不涉及
	Type *[]string `json:"type,omitempty"`

	// **参数解释**：当前IP地址关联的负载均衡实例域名。 支持多值查询，查询条件格式： *domain_name=xxx&domain_name=xxx*。  **约束限制**： - 如果IP为私网类型，则这里为负载均衡实例的私网域名。 - 如果IP为公网类型，则这里为负载均衡实例的公网域名。  **取值范围**：不涉及  **默认取值**：不涉及
	DomainName *[]string `json:"domain_name,omitempty"`

	// **参数解释**：所属的企业项目ID。 支持多值查询，查询条件格式： *enterprise_project_id=xxx&enterprise_project_id=xxx*。  **约束限制**： 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:loadbalancers:listDnsConfig权限。 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  **取值范围**：不涉及  **默认取值**：不涉及
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListDomainIPsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainIPsRequest struct{}"
	}

	return strings.Join([]string{"ListDomainIPsRequest", string(data)}, " ")
}
