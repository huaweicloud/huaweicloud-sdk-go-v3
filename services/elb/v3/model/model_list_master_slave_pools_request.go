package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMasterSlavePoolsRequest Request Object
type ListMasterSlavePoolsRequest struct {

	// **参数解释**：上一页最后一条记录的ID。  **约束限制**： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。  **取值范围**：不涉及  **默认取值**：不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：每页返回的个数。  **约束限制**：不涉及  **取值范围**：0-2000  **默认取值**：2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：是否反向查询。  **约束限制**： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。  **取值范围**： - true：查询上一页。 - false：查询下一页。  **默认取值**：false
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：后端服务器组的描述信息。 支持多值查询，查询条件格式：*description=xxx&description=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Description *[]string `json:"description,omitempty"`

	// **参数解释**：后端服务器组关联的健康检查的ID。 支持多值查询，查询条件格式：*healthmonitor_id=xxx&healthmonitor_id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	HealthmonitorId *[]string `json:"healthmonitor_id,omitempty"`

	// **参数解释**：后端服务器组的ID。 支持多值查询，查询条件格式：*id=xxx&id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Id *[]string `json:"id,omitempty"`

	// **参数解释**：后端服务器组的名称。 支持多值查询，查询条件格式：*name=xxx&name=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *[]string `json:"name,omitempty"`

	// **参数解释**：后端服务器组绑定的负载均衡器ID。 支持多值查询，查询条件格式：*loadbalancer_id=xxx&loadbalancer_id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`

	// **参数解释**：后端服务器组的后端协议。 支持多值查询，查询条件格式：*protocol=xxx&protocol=xxx*。  **约束限制**：不涉及  **取值范围**：TCP、UDP、IP、TLS、GRPC、HTTP、HTTPS和QUIC。  **默认取值**：不涉及
	Protocol *[]string `json:"protocol,omitempty"`

	// **参数解释**：后端服务器组的负载均衡算法。 支持多值查询，查询条件格式：*lb_algorithm=xxx&lb_algorithm=xxx*。  **约束限制**：不涉及  **取值范围**： 1、ROUND_ROBIN：加权轮询算法。 2、LEAST_CONNECTIONS：加权最少连接算法。 3、SOURCE_IP：源IP算法。 4、QUIC_CID：连接ID算法。  **默认取值**：不涉及
	LbAlgorithm *[]string `json:"lb_algorithm,omitempty"`

	// **参数解释**：资源所属的企业项目ID。 支持多值查询，查询条件格式：*enterprise_project_id=xxx&enterprise_project_id=xxx*。  **约束限制**： - 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:pools:list权限。 - 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  **取值范围**：不涉及  **默认取值**：不涉及  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：后端服务器组支持的IP版本。 支持多值查询，查询条件格式：*ip_version=xxx&ip_version=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	IpVersion *[]string `json:"ip_version,omitempty"`

	// **参数解释**：后端服务器的IP地址。 支持多值查询，查询条件格式：*member_address=xxx&member_address=xxx*。  **约束限制**：仅用于查询条件，不作为响应参数字段。  **取值范围**：不涉及  **默认取值**：不涉及
	MemberAddress *[]string `json:"member_address,omitempty"`

	// **参数解释**：后端服务器对应的弹性云服务器的ID。 支持多值查询，查询条件格式：*member_device_id=xxx&member_device_id=xxx*。  **约束限制**：仅用于查询条件，不作为响应参数字段。  **取值范围**：不涉及  **默认取值**：不涉及
	MemberDeviceId *[]string `json:"member_device_id,omitempty"`

	// **参数解释**：关联的监听器ID，包括通过l7policy关联的。 支持多值查询，查询条件格式：*listener_id=xxx&listener_id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ListenerId *[]string `json:"listener_id,omitempty"`

	// **参数解释**：后端服务器ID。 支持多值查询，查询条件格式：*member_instance_id=xxx&member_instance_id=xxx*。  **约束限制**：仅用于查询条件，不作为响应参数字段。  **取值范围**：不涉及  **默认取值**：不涉及
	MemberInstanceId *[]string `json:"member_instance_id,omitempty"`

	// **参数解释**：后端服务器组关联的虚拟私有云的ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	VpcId *[]string `json:"vpc_id,omitempty"`

	// **参数解释**：后端服务器组的类型。  **约束限制**：不涉及  **取值范围**： - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加IP类型后端，type指定为该类型时，vpc_id不允许指定。 - 空字符串（\"\"）：允许任意类型的后端  **默认取值**：不涉及
	Type *[]string `json:"type,omitempty"`

	// **参数解释**：查询是否开启延迟注销的功能，查询条件格式：*connection_drain=true或者*connection_drain=false  **约束限制**：不涉及  **取值范围**：true 开启，false 不开启。  **默认取值**：不涉及
	ConnectionDrain *bool `json:"connection_drain,omitempty"`
}

func (o ListMasterSlavePoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMasterSlavePoolsRequest struct{}"
	}

	return strings.Join([]string{"ListMasterSlavePoolsRequest", string(data)}, " ")
}
