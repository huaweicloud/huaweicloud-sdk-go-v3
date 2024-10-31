package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetFirewallInstanceResponseRecord struct {

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id。，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，默认情况下，fw_instance_Id为空时，返回账号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 防火墙名称
	Name *string `json:"name,omitempty"`

	// 集群类型，包含主备（0）和集群（1）两种方式，主备模式包含四个节点，2个主节点构成集群，剩余两个节点为主节点的备节点，集群模式仅拉起两个节点作为集群。
	HaType *int32 `json:"ha_type,omitempty"`

	// 计费模式 0：包年/包月 1：按需
	ChargeMode *int32 `json:"charge_mode,omitempty"`

	// 防火墙防护类型，目前仅支持0，互联网防护
	ServiceType *int32 `json:"service_type,omitempty"`

	// 引擎类型，0：自研引擎 1：山石引擎 3：天融信引擎
	EngineType *int32 `json:"engine_type,omitempty"`

	Flavor *Flavor `json:"flavor,omitempty"`

	// 防护对象列表
	ProtectObjects *[]ProtectObjectVo `json:"protect_objects,omitempty"`

	// 防火墙状态列表，包括-1：等待支付，0：创建中，1，删除中，2：运行中，3：升级中，4：删除完成：5：冻结中，6：创建失败，7：删除失败，8：冻结失败，9：存储中，10：存储失败，11：升级失败
	Status *int32 `json:"status,omitempty"`

	// 是否为旧引擎，true表示是，false表示不是
	IsOldFirewallInstance *bool `json:"is_old_firewall_instance,omitempty"`

	// 是否支持obs，true表示是，false表示不是
	IsAvailableObs *bool `json:"is_available_obs,omitempty"`

	// 是否支持威胁情报标签，true表示是，false表示不是
	IsSupportThreatTags *bool `json:"is_support_threat_tags,omitempty"`

	// 是否支持ipv6，true表示是，false表示不是
	SupportIpv6 *bool `json:"support_ipv6,omitempty"`

	// 特性开关，boolean值为true表示是，false表示否
	FeatureToggle map[string]bool `json:"feature_toggle,omitempty"`

	// 防火墙资源列表
	Resources *[]FirewallInstanceResource `json:"resources,omitempty"`

	// 防火墙名称
	FwInstanceName *string `json:"fw_instance_name,omitempty"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙资源id，同fw_instance_id
	ResourceId *string `json:"resource_id,omitempty"`

	// 是否支持url过滤，true表示是，false表示不是
	SupportUrlFiltering *bool `json:"support_url_filtering,omitempty"`

	// 标签列表，标签键值map转化的json字符串，如\"{\\\"key\\\":\\\"value\\\"}\"
	Tags *string `json:"tags,omitempty"`
}

func (o GetFirewallInstanceResponseRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetFirewallInstanceResponseRecord struct{}"
	}

	return strings.Join([]string{"GetFirewallInstanceResponseRecord", string(data)}, " ")
}
