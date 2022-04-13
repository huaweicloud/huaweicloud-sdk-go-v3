package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建云服务器组请求返回对象
type Pool struct {
	// 后端云服务器组的管理状态，只支持设置为true。  不支持该字段，请勿使用。

	AdminStateUp bool `json:"admin_state_up"`
	// 后端云服务器组的描述信息。

	Description string `json:"description"`
	// 后端云服务器组关联的健康检查的ID。

	HealthmonitorId string `json:"healthmonitor_id"`
	// 后端云服务器组的ID。

	Id string `json:"id"`
	// 后端云服务器组的负载均衡算法。  取值： 1、ROUND_ROBIN：加权轮询算法。 2、LEAST_CONNECTIONS：加权最少连接算法。 3、SOURCE_IP：源IP算法。 4、QUIC_CID：连接ID算法。  使用说明： - 当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。 - 只有pool的protocol为QUIC时，才支持QUIC_CID算法。

	LbAlgorithm string `json:"lb_algorithm"`
	// 后端云服务器组关联的监听器ID列表。实际上只会有一个关联的监听器ID。

	Listeners []ListenerRef `json:"listeners"`
	// 后端云服务器组关联的负载均衡器ID列表。实际只会有一个关联的负载均衡器ID。

	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`
	// 后端云服务器组中的后端云服务器ID列表。

	Members []MemberRef `json:"members"`
	// 后端云服务器组的名称。

	Name string `json:"name"`
	// 后端云服务器组所在的项目ID。

	ProjectId string `json:"project_id"`
	// 后端云服务器组的后端协议。  取值：TCP、UDP、HTTP、HTTPS和QUIC。  使用说明： - listener的protocol为UDP时，pool的protocol必须为UDP或QUIC； - listener的protocol为TCP时pool的protocol必须为TCP； - listener的protocol为HTTP时，pool的protocol必须为HTTP。 - listener的protocol为HTTPS时，pool的protocol必须为HTTP或HTTPS。 - listener的protocol为TERMINATED_HTTPS时，pool的protocol必须为HTTP。

	Protocol string `json:"protocol"`

	SessionPersistence *SessionPersistence `json:"session_persistence"`
	// 后端云服务器组支持的IP版本。[取值：  [共享型：默认为v4；](tag:hc,hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42,dt,dt_test)  [独享型：取值范围(dualstack、v4、v6)。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。当协议为HTTP时，ip_version为v4。](tag:hc,hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42,dt,dt_test) [取值范围(dualstack、v4、v6)。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。当协议为HTTP时，ip_version为v4。](tag:hcso_dt)  [不支持IPv6，只会返回v4。](tag:dt,dt_test)

	IpVersion string `json:"ip_version"`

	SlowStart *SlowStart `json:"slow_start"`
	// 是否开启误删保护。取值：false不开启，true开启。  > 退场时需要先关闭所有资源的删除保护开关。

	MemberDeletionProtectionEnable bool `json:"member_deletion_protection_enable"`
}

func (o Pool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pool struct{}"
	}

	return strings.Join([]string{"Pool", string(data)}, " ")
}
