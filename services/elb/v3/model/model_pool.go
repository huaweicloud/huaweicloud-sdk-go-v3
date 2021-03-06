package model

import (
	"encoding/json"

	"strings"
)

// 创建云服务器组请求返回对象
type Pool struct {
	// 后端云服务器组的管理状态；该字段为预留字段，暂未启用。只支持更新为true。

	AdminStateUp bool `json:"admin_state_up"`
	// 后端云服务器组的描述信息

	Description string `json:"description"`
	// 后端云服务器组关联的健康检查的ID。

	HealthmonitorId string `json:"healthmonitor_id"`
	// 后端云服务器组的ID。

	Id string `json:"id"`
	// 后端云服务器组的负载均衡算法，取值：ROUND_ROBIN：加权轮询算法；LEAST_CONNECTIONS：加权最少连接算法；SOURCE_IP：源IP算法；当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。

	LbAlgorithm string `json:"lb_algorithm"`
	// 后端云服务器组关联的监听器ID的列表。

	Listeners []ListenerRef `json:"listeners"`
	// 后端云服务器组绑定的负载均衡器ID的列表。

	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`
	// 后端云服务器组关联的后端云服务器ID的列表。

	Members []MemberRef `json:"members"`
	// 后端云服务器组的名称。

	Name string `json:"name"`
	// 后端云服务器组所在的项目ID。

	ProjectId string `json:"project_id"`
	// 后端云服务器组的后端协议。 使用说明：支持TCP、UDP和HTTP。listener的protocol为UDP时pool的protocol必须为UDP；listener的protocol为TCP时pool的protocol必须为TCP；listener的protocol为HTTP或TERMINATED_HTTPS时pool的protocol必须为HTTP。

	Protocol string `json:"protocol"`

	SessionPersistence *SessionPersistence `json:"session_persistence"`
	// 后端云服务器组支持的IP版本， 共享型：默认为v4； 性能保障型：取值范围(dualstack、v4、v6)。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。当协议为HTTP时，ip_version为v4。

	IpVersion string `json:"ip_version"`

	SlowStart *SlowStart `json:"slow_start"`
	// 是否开启误删保护，默认false不开启

	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`
}

func (o Pool) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Pool struct{}"
	}

	return strings.Join([]string{"Pool", string(data)}, " ")
}
