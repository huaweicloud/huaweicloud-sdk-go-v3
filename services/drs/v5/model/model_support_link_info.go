package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupportLinkInfo 支持的链路信息
type SupportLinkInfo struct {

	// **参数解释**：  DRS任务引擎类型，标识数据复制任务所使用的数据库引擎。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	EngineType *string `json:"engine_type,omitempty"`

	// **参数解释**：  任务使用的网络类型。  **约束限制**：  灾备场景不支持选择VPC网络。  **取值范围**：  - eip：公网网络 - vpc：VPC网络 - vpn：VPN、专线网络  **默认取值**：  不涉及
	NetType *string `json:"net_type,omitempty"`

	// **参数解释**：  可用链路支持的迁移模式列表，表示该链路支持的数据迁移执行方式。  **约束限制**：  不涉及。  **取值范围**：  - FULL_TRANS：全量迁移 - FULL_INCR_TRANS：全量+增量迁移 - INCR_TRANS：增量迁移  **默认取值**：  不涉及
	TaskModes *[]string `json:"task_modes,omitempty"`

	// **参数解释**：  任务的迁移方向。  **约束限制**：  不涉及。  **取值范围**：  - up：入云，灾备场景时对应本云为备 - down：出云，灾备场景时对应本云为主 - non-dbs：自建  **默认取值**：  不涉及
	JobDirection *string `json:"job_direction,omitempty"`

	// **参数解释**：  云上数据库实例的部署模式，标识目标或源数据库实例在云上的集群架构类型。  **约束限制**：  不涉及。  **取值范围**：  - Single：单机模式 - Ha：主备模式 - Cluster：集群模式 - Sharding：分片模式 - Independent：GaussDB独立部署模式  **默认取值**：  不涉及
	ClusterMode *string `json:"cluster_mode,omitempty"`

	// **参数解释**：  DRS迁移实例的部署类型，决定迁移任务所使用的实例架构。  **约束限制**：  不涉及。  **取值范围**：  - Single：单机 - Ha：主备  **默认取值**：  不涉及
	JobInstanceType *string `json:"job_instance_type,omitempty"`

	// **参数解释**：  是否支持绑定弹性公网IP（EIP），标识当前链路是否允许为迁移实例绑定弹性公网IP以实现公网通信。  **约束限制**：  不涉及。  **取值范围**：  - true：支持绑定EIP - false：不支持绑定EIP  **默认取值**：  不涉及
	IsSupportBindEip *bool `json:"is_support_bind_eip,omitempty"`
}

func (o SupportLinkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportLinkInfo struct{}"
	}

	return strings.Join([]string{"SupportLinkInfo", string(data)}, " ")
}
