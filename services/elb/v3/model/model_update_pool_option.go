package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新后端服务器组请求参数。
type UpdatePoolOption struct {

	// 后端云服务器组的管理状态，只支持更新为true。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 后端云服务器组的描述信息。
	Description *string `json:"description,omitempty"`

	// 后端云服务器组的负载均衡算法。  取值： - ROUND_ROBIN：加权轮询算法。 - LEAST_CONNECTIONS：加权最少连接算法。 - SOURCE_IP：源IP算法。 - QUIC_CID：连接ID算法。  使用说明： - 当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。 - 只有pool的protocol为QUIC时，才支持QUIC_CID算法。  [不支持QUIC_CID算法。](tag:hws_eu,g42,hk_g42,hcso_dt)  [荷兰region不支持QUIC。](tag:dt)
	LbAlgorithm *string `json:"lb_algorithm,omitempty"`

	// 后端云服务器组的名称。
	Name *string `json:"name,omitempty"`

	SessionPersistence *UpdatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	SlowStart *UpdatePoolSlowStartOption `json:"slow_start,omitempty"`

	// 是否开启删除保护。  取值：false不开启，true开启。  > 退场时需要先关闭所有资源的删除保护开关。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42)  [荷兰region不支持该字段，请勿使用。](tag:dt)
	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	// 后端云服务器组关联的虚拟私有云的ID。  只有vpc_id为空时允许更新。
	VpcId *string `json:"vpc_id,omitempty"`

	// 后端服务器组的类型。  取值： - instance：允许任意类型的后端，type指定为该类型时，vpc_id是必选字段。 - ip：只能添加跨VPC后端，type指定为该类型时，vpc_id不允许指定。 - 空字符串（\"\"）：允许任意类型的后端  使用说明： - 只有type为空时允许更新，不允许从非空更新为空。
	Type *string `json:"type,omitempty"`
}

func (o UpdatePoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolOption struct{}"
	}

	return strings.Join([]string{"UpdatePoolOption", string(data)}, " ")
}
