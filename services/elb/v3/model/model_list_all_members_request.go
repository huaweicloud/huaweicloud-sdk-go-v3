package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllMembersRequest Request Object
type ListAllMembersRequest struct {

	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。
	Marker *string `json:"marker,omitempty"`

	// 参数解释：每页返回的个数。  取值范围：0-2000  默认取值：2000
	Limit *int32 `json:"limit,omitempty"`

	// 是否反向查询。  取值： - true：查询上一页。 - false：查询下一页，默认。  使用说明： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 后端服务器名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。
	Name *[]string `json:"name,omitempty"`

	// 后端服务器的权重，请求按权重在同一后端服务器组下的后端服务器间分发。 权重为0的后端不再接受新的请求。 当后端服务器所在的后端服务器组的lb_algorithm的取值为SOURCE_IP时，该字段无效。  支持多值查询，查询条件格式：*weight=xxx&weight=xxx*。
	Weight *[]int32 `json:"weight,omitempty"`

	// 后端服务器的管理状态；该字段虽然支持创建、更新，但实际取值决定于member对应的弹性云服务器是否存在。 若存在，该值为true，否则，该值为false。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 后端服务器所在的子网ID。该子网和后端服务器关联的负载均衡器的子网必须在同一VPC下。只支持指定IPv4的子网ID。  支持多值查询，查询条件格式：***subnet_cidr_id=xxx&subnet_cidr_id=xxx*。
	SubnetCidrId *[]string `json:"subnet_cidr_id,omitempty"`

	// 后端服务器的对应的IP地址，这个IP必须在subnet_cidr_id字段的子网网段中。 例如：192.168.3.11。  支持多值查询，查询条件格式：*address=xxx&address=xxx*。
	Address *[]string `json:"address,omitempty"`

	// 后端服务器端口号。  支持多值查询，查询条件格式：*protocol_port=xxx&protocol_port=xxx*。
	ProtocolPort *[]int32 `json:"protocol_port,omitempty"`

	// 后端服务器ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。
	Id *[]string `json:"id,omitempty"`

	// 后端服务器的健康状态。  取值： - INITIAL：初始化中，表示负载均衡实例配置了健康检查，但查不到数据。 - ONLINE，后端服务器正常运行。 - NO_MONITOR，后端服务器无健康检查。 - OFFLINE，已下线。  支持多值查询，查询条件格式：*operating_status=xxx&operating_status=*。
	OperatingStatus *[]string `json:"operating_status,omitempty"`

	// 参数解释：所属的企业项目ID。 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:members:list权限。 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  支持多值查询，查询条件格式： *enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// IP版本，取值v4、v6。  支持多值查询，查询条件格式：*ip_version=xxx&ip_version=xxx*。
	IpVersion *[]string `json:"ip_version,omitempty"`

	// member所属的服务器组ID  支持多值查询，查询条件格式：*pool_id=xxx&pool_id=xxx*。
	PoolId *[]string `json:"pool_id,omitempty"`

	// member所属的负载均衡器ID。  支持多值查询，查询条件格式：*loadbalancer_id=xxx&loadbalancer_id=xxx*。
	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`
}

func (o ListAllMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllMembersRequest struct{}"
	}

	return strings.Join([]string{"ListAllMembersRequest", string(data)}, " ")
}
