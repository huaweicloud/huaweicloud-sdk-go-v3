package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPoolsRequest struct {
	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。使用说明：必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 后端云服务器组的描述信息。  支持多值查询，查询条件格式：*description=xxx&description=xxx*。

	Description *[]string `json:"description,omitempty"`
	// 后端云服务器组的管理状态。  不支持该字段，请勿使用。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 后端云服务器组关联的健康检查的ID。  支持多值查询，查询条件格式：*healthmonitor_id=xxx&healthmonitor_id=xxx*。

	HealthmonitorId *[]string `json:"healthmonitor_id,omitempty"`
	// 后端云服务器组的ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。

	Id *[]string `json:"id,omitempty"`
	// 后端云服务器组的名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。

	Name *[]string `json:"name,omitempty"`
	// 后端云服务器组绑定的负载均衡器ID。  支持多值查询，查询条件格式：*loadbalancer_id=xxx&loadbalancer_id=xxx*。

	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`
	// 后端云服务器组的后端协议。取值：TCP、UDP、HTTP、HTTPS和QUIC。  支持多值查询，查询条件格式：*protocol=xxx&protocol=xxx*。

	Protocol *[]string `json:"protocol,omitempty"`
	// 后端云服务器组的负载均衡算法。  取值： 1、ROUND_ROBIN：加权轮询算法。 2、LEAST_CONNECTIONS：加权最少连接算法。 3、SOURCE_IP：源IP算法。 4、QUIC_CID：连接ID算法。  支持多值查询，查询条件格式：*lb_algorithm=xxx&lb_algorithm=xxx*。

	LbAlgorithm *[]string `json:"lb_algorithm,omitempty"`
	// 企业项目ID。不传时查询default企业项目\"0\"下的资源，鉴权按照default企业项目鉴权；如果传值，则传已存在的企业项目ID或all_granted_eps（表示查询所有企业项目）进行查询。 支持多值查询，查询条件格式：*enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// 后端云服务器组支持的IP版本。取值： - 共享型LB下的pool：固定为v4； -   独享型LB下的pool：dualstack、v4。当该pool的协议为TCP/UDP/QUIC时，ip_version为dualstack，表示双栈。当协议为HTTP/HTTPS时，ip_version为v4。 [不支持IPv6，只会返回v4。](tag:otc,otc_test,dt,dt_test) 支持多值查询，查询条件格式：*ip_version=xxx&ip_version=xxx*。

	IpVersion *[]string `json:"ip_version,omitempty"`
	// 后端云服务器的IP地址。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_address=xxx&member_address=xxx*。

	MemberAddress *[]string `json:"member_address,omitempty"`
	// 后端云服务器对应的弹性云服务器的ID。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_device_id=xxx&member_device_id=xxx*。

	MemberDeviceId *[]string `json:"member_device_id,omitempty"`
	// 是否开启删除保护，false不开启，true开启，不传查询全部。

	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`
	// 关联的监听器ID，包括通过l7policy关联的。  支持多值查询，查询条件格式：*listener_id=xxx&listener_id=xxx*。

	ListenerId *[]string `json:"listener_id,omitempty"`
	// 后端云服务器ID。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_instance_id=xxx&member_instance_id=xxx*。

	MemberInstanceId *[]string `json:"member_instance_id,omitempty"`
}

func (o ListPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListPoolsRequest", string(data)}, " ")
}
