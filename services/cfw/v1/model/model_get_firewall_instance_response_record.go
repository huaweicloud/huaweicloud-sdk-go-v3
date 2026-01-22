package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetFirewallInstanceResponseRecord struct {

	// **参数解释**： 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id。可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获取 **取值范围**： 不涉及
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// **参数解释**： 防火墙名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 集群类型 **取值范围**： - 0：主备模式，包含四个节点，2个主节点构成集群，剩余两个节点为主节点的备节点 - 1：集群模式，仅拉起两个节点作为集群
	HaType *int32 `json:"ha_type,omitempty"`

	// **参数解释**： 计费模式 **取值范围**： - 0：包年/包月 - 1：按需
	ChargeMode *int32 `json:"charge_mode,omitempty"`

	// **参数解释**： 防火墙防护类型 **取值范围**： 目前仅支持0，互联网防护
	ServiceType *int32 `json:"service_type,omitempty"`

	// **参数解释**： 引擎类型 **取值范围**： - 0：自研引擎 - 1：山石引擎 - 3：天融信引擎
	EngineType *int32 `json:"engine_type,omitempty"`

	Flavor *Flavor `json:"flavor,omitempty"`

	// **参数解释**： 防护对象列表 **取值范围**： 不涉及
	ProtectObjects *[]ProtectObjectVo `json:"protect_objects,omitempty"`

	// **参数解释**： 防火墙状态列表 **取值范围**： - -1：等待支付 - 0：创建中 - 1，删除中 - 2：运行中 - 3：升级中 - 4：删除完成 - 5：冻结中 - 6：创建失败 - 7：删除失败 - 8：冻结失败 - 9：存储中 - 10：存储失败 - 11：升级失败
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 是否为旧引擎 **取值范围**： - true：是 - false：不是
	IsOldFirewallInstance *bool `json:"is_old_firewall_instance,omitempty"`

	// **参数解释**： 是否支持obs **取值范围**： - true：是 - false：不是
	IsAvailableObs *bool `json:"is_available_obs,omitempty"`

	// **参数解释**： 是否支持威胁情报标签 **取值范围**： - true：是 - false：不是
	IsSupportThreatTags *bool `json:"is_support_threat_tags,omitempty"`

	// **参数解释**： 是否支持ipv6 **取值范围**： - true：是 - false：不是
	SupportIpv6 *bool `json:"support_ipv6,omitempty"`

	// **参数解释**： 特性开关 **取值范围**： - true：是 - false：不是
	FeatureToggle map[string]bool `json:"feature_toggle,omitempty"`

	// **参数解释**： 防火墙资源列表 **取值范围**： 不涉及
	Resources *[]FirewallInstanceResource `json:"resources,omitempty"`

	// **参数解释**： 防火墙名称 **取值范围**： 不涉及
	FwInstanceName *string `json:"fw_instance_name,omitempty"`

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取 **取值范围**： 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 防火墙资源id，同fw_instance_id **取值范围**： 不涉及
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 是否支持url过滤 **取值范围**： - true：是 - false：不是
	SupportUrlFiltering *bool `json:"support_url_filtering,omitempty"`

	// **参数解释**： 标签列表，标签键值map转化的json字符串，如\"{\\\"key\\\":\\\"value\\\"}\" **取值范围**： 不涉及
	Tags *string `json:"tags,omitempty"`
}

func (o GetFirewallInstanceResponseRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetFirewallInstanceResponseRecord struct{}"
	}

	return strings.Join([]string{"GetFirewallInstanceResponseRecord", string(data)}, " ")
}
