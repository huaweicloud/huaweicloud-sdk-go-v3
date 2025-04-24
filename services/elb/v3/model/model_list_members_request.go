package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMembersRequest Request Object
type ListMembersRequest struct {

	// 参数解释：后端服务器组ID。
	PoolId string `json:"pool_id"`

	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。
	Marker *string `json:"marker,omitempty"`

	// 参数解释：每页返回的个数。  取值范围：0-2000  默认取值：2000
	Limit *int32 `json:"limit,omitempty"`

	// 是否反向查询。  取值： - true：查询上一页。 - false：查询下一页，默认。  使用说明： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 后端服务器名称。注意：该名称并非ECS名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。
	Name *[]string `json:"name,omitempty"`

	// 后端服务器的权重，请求将根据pool配置的负载均衡算法和后端服务器的权重进行负载分发。 权重值越大，分发的请求越多。权重为0的后端不再接受新的请求。  取值：0-100。  支持多值查询，查询条件格式：*weight=xxx&weight=xxx*。
	Weight *[]int32 `json:"weight,omitempty"`

	// 后端服务器的管理状态。  取值：true、false。  虽然创建、更新请求支持该字段，但实际取值决定于后端服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 后端服务器所在子网的IPv4子网ID或IPv6子网ID。  支持多值查询，查询条件格式：***subnet_cidr_id=xxx&subnet_cidr_id=xxx*。  [不支持IPv6，请勿设置为IPv6子网ID。](tag:dt)
	SubnetCidrId *[]string `json:"subnet_cidr_id,omitempty"`

	// 后端服务器对应的IPv4或IPv6地址。  支持多值查询，查询条件格式：*address=xxx&address=xxx*。  [不支持IPv6，请勿设置为IPv6地址。](tag:dt)
	Address *[]string `json:"address,omitempty"`

	// 后端服务器业务端口号。  支持多值查询，查询条件格式：*protocol_port=xxx&protocol_port=xxx*。
	ProtocolPort *[]int32 `json:"protocol_port,omitempty"`

	// 后端服务器ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。
	Id *[]string `json:"id,omitempty"`

	// 后端服务器的健康状态。  取值： - ONLINE：后端服务器正常。 - NO_MONITOR：后端服务器所在的服务器组没有健康检查器。 - OFFLINE：后端服务器关联的ECS服务器不存在或已关机。  支持多值查询，查询条件格式：*operating_status=xxx&operating_status=xxx*。
	OperatingStatus *[]string `json:"operating_status,omitempty"`

	// 参数解释：所属的企业项目ID。 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:members:list权限。 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  支持多值查询，查询条件格式： *enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 当前后端服务器的IP地址版本。取值：v4、v6。
	IpVersion *[]string `json:"ip_version,omitempty"`

	// 后端服务器的类型。  取值： - ip：跨VPC的member。 - instance：关联到ECS的member。  支持多值查询，查询条件格式：*member_type=xxx&member_type=xxx*。
	MemberType *[]string `json:"member_type,omitempty"`

	// member关联的ECS实例ID，空表示跨VPC场景的member。  支持多值查询，查询条件格式：*instance_id=xxx&instance_id=xxx*。
	InstanceId *[]string `json:"instance_id,omitempty"`
}

func (o ListMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMembersRequest struct{}"
	}

	return strings.Join([]string{"ListMembersRequest", string(data)}, " ")
}
