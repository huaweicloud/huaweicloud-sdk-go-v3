package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FirewallInstanceVo struct {

	// **参数解释**： 防火墙实例id，创建云防火墙后用于标记防火墙由系统自动生成的id。 **约束限制**： 不涉及
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// **参数解释**： 资源id，与防火墙实例id fw_instance_id相同。 **约束限制**： 不涉及
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 创建防火墙时的时间戳 **约束限制**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 防火墙名称 **约束限制**： 不涉及
	FwInstanceName *string `json:"fw_instance_name,omitempty"`

	// **参数解释**： 企业项目id，用户支持企业项目后，由企业项目生成的id。 **约束限制**： 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 集群类型 **约束限制**： 不涉及 **取值范围**： - 0：主备模式，包含四个节点，2个主节点构成集群，剩余两个节点为主节点的备节点 - 1：集群模式，仅拉起两个节点作为集群
	HaType *int32 `json:"ha_type,omitempty"`

	// **参数解释**： 计费模式 **约束限制**： 不涉及 **取值范围**： - 0：包年/包月 - 1：按需 **默认取值**： 不涉及
	ChargeMode *int32 `json:"charge_mode,omitempty"`

	// **参数解释**： 防火墙防护类型 **约束限制**： 不涉及 **取值范围**： 目前仅支持0，互联网防护
	ServiceType *int32 `json:"service_type,omitempty"`

	// **参数解释**： 引擎类型 **约束限制**： 不涉及 **取值范围**： - 0：自研引擎 - 1：山石引擎 - 3：天融信引擎
	EngineType *int32 `json:"engine_type,omitempty"`

	Flavor *Flavor `json:"flavor,omitempty"`

	// **参数解释**： 防火墙状态列表 **约束限制**： 不涉及 **取值范围**： - -1：等待支付 - 0：创建中 - 1，删除中 - 2：运行中 - 3：升级中 - 4：删除完成 - 5：冻结中 - 6：创建失败 - 7：删除失败 - 8：冻结失败 - 9：存储中 - 10：存储失败 - 11：升级失败
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 标签列表，标签键值map转化的json字符串，如\"{\\\"key\\\":\\\"value\\\"}\" **约束限制**： 不涉及
	Tags *string `json:"tags,omitempty"`
}

func (o FirewallInstanceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirewallInstanceVo struct{}"
	}

	return strings.Join([]string{"FirewallInstanceVo", string(data)}, " ")
}
