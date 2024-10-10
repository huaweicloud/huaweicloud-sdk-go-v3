package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMasterSlavePoolsRequest Request Object
type ListMasterSlavePoolsRequest struct {

	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。
	Marker *string `json:"marker,omitempty"`

	// 参数解释：每页返回的个数。  取值范围：0-2000  默认取值：2000
	Limit *int32 `json:"limit,omitempty"`

	// 是否反向查询。  取值： - true：查询上一页。 - false：查询下一页，默认。  使用说明： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 后端服务器组的描述信息。  支持多值查询，查询条件格式：*description=xxx&description=xxx*。
	Description *[]string `json:"description,omitempty"`

	// 后端服务器组关联的健康检查的ID。  支持多值查询，查询条件格式：*healthmonitor_id=xxx&healthmonitor_id=xxx*。
	HealthmonitorId *[]string `json:"healthmonitor_id,omitempty"`

	// 后端服务器组的ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。
	Id *[]string `json:"id,omitempty"`

	// 后端服务器组的名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。
	Name *[]string `json:"name,omitempty"`

	// 后端服务器组绑定的负载均衡器ID。  支持多值查询，查询条件格式：*loadbalancer_id=xxx&loadbalancer_id=xxx*。
	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`

	// 后端服务器组的后端协议。  取值：TCP、UDP、[IP、](tag:hws_eu)TLS、GRPC、HTTP、HTTPS和QUIC。 [IP类型为网关型LB独有的后端服务器组协议。](tag:hws_eu)  支持多值查询，查询条件格式：*protocol=xxx&protocol=xxx*。  [不支持QUIC。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt,dt_test)
	Protocol *[]string `json:"protocol,omitempty"`

	// 后端服务器组的负载均衡算法。  取值： 1、ROUND_ROBIN：加权轮询算法。 2、LEAST_CONNECTIONS：加权最少连接算法。 3、SOURCE_IP：源IP算法。 4、QUIC_CID：连接ID算法。  支持多值查询，查询条件格式：*lb_algorithm=xxx&lb_algorithm=xxx*。  [不支持QUIC_CID。](tag:tm,hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC_CID。](tag:dt,dt_test)
	LbAlgorithm *[]string `json:"lb_algorithm,omitempty"`

	// 企业项目ID。不传时查询default企业项目\"0\"下的资源，鉴权按照default企业项目鉴权； 如果传值，则传已存在的企业项目ID或all_granted_eps（表示查询所有企业项目）进行查询。  支持多值查询，查询条件格式： *enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 后端服务器组支持的IP版本。  支持多值查询，查询条件格式：*ip_version=xxx&ip_version=xxx*。
	IpVersion *[]string `json:"ip_version,omitempty"`

	// 后端服务器的IP地址。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_address=xxx&member_address=xxx*。
	MemberAddress *[]string `json:"member_address,omitempty"`

	// 后端服务器对应的弹性云服务器的ID。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_device_id=xxx&member_device_id=xxx*。
	MemberDeviceId *[]string `json:"member_device_id,omitempty"`

	// 关联的监听器ID，包括通过l7policy关联的。  支持多值查询，查询条件格式：*listener_id=xxx&listener_id=xxx*。
	ListenerId *[]string `json:"listener_id,omitempty"`

	// 后端服务器ID。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_instance_id=xxx&member_instance_id=xxx*。
	MemberInstanceId *[]string `json:"member_instance_id,omitempty"`

	// 后端服务器组关联的虚拟私有云的ID。
	VpcId *[]string `json:"vpc_id,omitempty"`

	// 后端服务器组的类型。  取值： - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加跨VPC后端，type指定为该类型时，vpc_id不允许指定。 - 空字符串（\"\"）：允许任意类型的后端
	Type *[]string `json:"type,omitempty"`

	// 查询是否开启延迟注销的功能，查询条件格式：*connection_drain=true或者*connection_drain=false
	ConnectionDrain *bool `json:"connection_drain,omitempty"`
}

func (o ListMasterSlavePoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMasterSlavePoolsRequest struct{}"
	}

	return strings.Join([]string{"ListMasterSlavePoolsRequest", string(data)}, " ")
}
