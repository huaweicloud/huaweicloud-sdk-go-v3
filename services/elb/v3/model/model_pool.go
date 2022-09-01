package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建云服务器组请求返回对象
type Pool struct {

	// 后端云服务器组的管理状态，只支持设置为true。  不支持该字段，请勿使用。
	AdminStateUp bool `json:"admin_state_up" xml:"admin_state_up"`

	// 后端云服务器组的描述信息。
	Description string `json:"description" xml:"description"`

	// 后端云服务器组关联的健康检查的ID。
	HealthmonitorId string `json:"healthmonitor_id" xml:"healthmonitor_id"`

	// 后端云服务器组的ID。
	Id string `json:"id" xml:"id"`

	// 后端云服务器组的负载均衡算法。  取值： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。  使用说明： - 当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。 - 只有pool的protocol为QUIC时，才支持QUIC_CID算法。
	LbAlgorithm string `json:"lb_algorithm" xml:"lb_algorithm"`

	// 后端云服务器组关联的监听器ID列表。
	Listeners []ListenerRef `json:"listeners" xml:"listeners"`

	// 后端云服务器组关联的负载均衡器ID列表。
	Loadbalancers []LoadBalancerRef `json:"loadbalancers" xml:"loadbalancers"`

	// 后端云服务器组中的后端云服务器ID列表。
	Members []MemberRef `json:"members" xml:"members"`

	// 后端云服务器组的名称。
	Name string `json:"name" xml:"name"`

	// 后端云服务器组所在的项目ID。
	ProjectId string `json:"project_id" xml:"project_id"`

	// 后端云服务器组的后端协议。  取值：TCP、UDP、HTTP、HTTPS和QUIC。   使用说明： - listener的protocol为UDP时，pool的protocol必须为UDP或QUIC； - listener的protocol为TCP时pool的protocol必须为TCP； - listener的protocol为HTTP时，pool的protocol必须为HTTP。 - listener的protocol为HTTPS时，pool的protocol必须为HTTP或HTTPS。 - listener的protocol为TERMINATED_HTTPS时，pool的protocol必须为HTTP。
	Protocol string `json:"protocol" xml:"protocol"`

	SessionPersistence *SessionPersistence `json:"session_persistence" xml:"session_persistence"`

	// 后端云服务器组支持的IP版本。 [取值： - 共享型：固定为v4； - 独享型：取值dualstack、v4、v6。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。当协议为HTTP时，ip_version为v4。](tag:hws,hws_hk,ocb,tlf,ctc,hcs,sbc,g42,tm,cmcc,hk_g42,mix,hk_sbc,hws_ocb,fcs) [取值：dualstack、v4、v6。当协议为TCP/UDP时，ip_version为dualstack，表示双栈。当协议为HTTP时，ip_version为v4。](tag:hcso_dt)   [不支持IPv6，只会返回v4。](tag:dt,dt_test)
	IpVersion string `json:"ip_version" xml:"ip_version"`

	SlowStart *SlowStart `json:"slow_start" xml:"slow_start"`

	// 是否开启误删保护。取值：false不开启，true开启。 > 退场时需要先关闭所有资源的删除保护开关。
	MemberDeletionProtectionEnable bool `json:"member_deletion_protection_enable" xml:"member_deletion_protection_enable"`

	// 创建时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 后端云服务器组关联的虚拟私有云的ID。
	VpcId string `json:"vpc_id" xml:"vpc_id"`

	// 后端服务器组的类型。   取值：  - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。  - ip：只能添加跨VPC后端，type指定为该类型时，vpc_id不允许指定。  - 空字符串（\"\"）：允许任意类型的后端
	Type string `json:"type" xml:"type"`
}

func (o Pool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pool struct{}"
	}

	return strings.Join([]string{"Pool", string(data)}, " ")
}
