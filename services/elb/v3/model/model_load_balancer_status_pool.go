package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoadBalancerStatusPool LB状态树的后端服务器组状态信息。
type LoadBalancerStatusPool struct {

	// 后端服务器组的配置状态。  取值： - ACTIVE：使用中。
	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	// 后端服务器组名。
	Name *string `json:"name,omitempty"`

	Healthmonitor *LoadBalancerStatusHealthMonitor `json:"healthmonitor,omitempty"`

	// 后端服务器状态信息。
	Members *[]LoadBalancerStatusMember `json:"members,omitempty"`

	// 参数解释：后端服务器组ID。
	Id *string `json:"id,omitempty"`

	// 后端服务器组的操作状态。  取值： - ONLINE：创建时默认状态，表后端服务器组正常。 - DEGRADED：该后端服务器组下存在member为的operating_status=OFFLINE。 - DISABLED：负载均衡器或后端服务器组的admin_state_up=false。  说明： DEGRADED和DISABLED仅在当前接口返回， 查询后端服务器组详情等其他接口返回的operating_status字段不存在这两个状态值。
	OperatingStatus *string `json:"operating_status,omitempty"`
}

func (o LoadBalancerStatusPool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusPool struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusPool", string(data)}, " ")
}
