package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceTopologyResponse Response Object
type ShowInstanceTopologyResponse struct {

	// **参数解释**： 集群角色：redis-server实例拓朴图。
	RedisServer *[]TopologyInfo `json:"redis_server,omitempty"`

	// **参数解释**： 集群角色：lvs实例拓朴图（当前已下线），只适用于Redis 3.0版本实例。
	ClusterLvs *[]TopologyInfo `json:"cluster_lvs,omitempty"`

	// **参数解释**： 集群角色：admin实例拓朴图（当前已下线），只适用于Redis 3.0版本实例。
	ClusterAdmin *[]TopologyInfo `json:"cluster_admin,omitempty"`

	// **参数解释**： 集群角色：proxy实例拓朴图。
	ClusterProxy *[]TopologyInfo `json:"cluster_proxy,omitempty"`

	// **参数解释**： Redis的角色master实例拓朴图，只适用于Redis 3.0版本实例。
	Master *[]TopologyInfo `json:"master,omitempty"`

	// **参数解释**： 集群角色：VPC Endpoint服务实例拓朴图，适用于4.0及以上版本实例。
	Vpcendpoint *[]TopologyInfo `json:"vpcendpoint,omitempty"`

	// **参数解释**： 集群角色：ELB服务实例拓朴图，适用于Redis 4.0及以上版本实例。
	Elb            *[]TopologyInfo `json:"elb,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowInstanceTopologyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTopologyResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceTopologyResponse", string(data)}, " ")
}
