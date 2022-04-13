package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建主机组请求
type CreatePoolOption struct {
	// 后端云服务器组的管理状态，只支持更新为true。  不支持该字段，请勿使用。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 后端云服务器组的描述信息。

	Description *string `json:"description,omitempty"`
	// 后端云服务器组的负载均衡算法。  取值： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。  使用说明： - 当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。 - 只有pool的protocol为QUIC时，才支持QUIC_CID算法。

	LbAlgorithm string `json:"lb_algorithm"`
	// 后端云服务器组关联的监听器的ID。  使用说明： - listener_id和loadbalancer_id至少指定一个。

	ListenerId *string `json:"listener_id,omitempty"`
	// 后端云服务器组关联的负载均衡器ID。  使用说明： - listener_id和loadbalancer_id中至少指定一个。

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
	// 后端云服务器组的名称。

	Name *string `json:"name,omitempty"`
	// 后端云服务器组所属的项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 后端云服务器组的后端协议。  取值：TCP、UDP、HTTP、HTTPS和QUIC。  使用说明： - listener的protocol为UDP时，pool的protocol必须为UDP或QUIC； - listener的protocol为TCP时pool的protocol必须为TCP； - listener的protocol为HTTP时，pool的protocol必须为HTTP。 - listener的protocol为HTTPS时，pool的protocol必须为HTTP或HTTPS。 - listener的protocol为TERMINATED_HTTPS时，pool的protocol必须为HTTP。

	Protocol string `json:"protocol"`

	SessionPersistence *CreatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	SlowStart *CreatePoolSlowStartOption `json:"slow_start,omitempty"`
	// 是否开启删除保护。取值：false不开启，true开启，默认false。 > 退场时需要先关闭所有资源的删除保护开关。

	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`
}

func (o CreatePoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolOption struct{}"
	}

	return strings.Join([]string{"CreatePoolOption", string(data)}, " ")
}
