package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FirewallInstanceVo struct {

	// 防火墙实例id，创建云防火墙后用于标记防火墙由系统自动生成的id。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 资源id，与防火墙实例id fw_instance_id相同
	ResourceId *string `json:"resource_id,omitempty"`

	// 创建防火墙时的时间戳
	Name *string `json:"name,omitempty"`

	// 防火墙名称
	FwInstanceName *string `json:"fw_instance_name,omitempty"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 集群类型，包含主备（0）和集群（1）两种方式，主备模式包含四个节点，2个主节点构成集群，剩余两个节点为主节点的备节点，集群模式仅拉起两个节点作为集群。
	HaType *int32 `json:"ha_type,omitempty"`

	// 计费模式 0：包年/包月 1：按需
	ChargeMode *int32 `json:"charge_mode,omitempty"`

	// 防火墙防护类型，目前仅支持0，互联网防护。
	ServiceType *int32 `json:"service_type,omitempty"`

	// 引擎类型，0：自研引擎 1：山石引擎 3：天融信引擎
	EngineType *int32 `json:"engine_type,omitempty"`

	Flavor *Flavor `json:"flavor,omitempty"`

	// 防火墙状态列表，包括-1：等待支付，0：创建中，1：删除中，2：运行中，3：升级中，4：删除完成：5：冻结中，6：创建失败，7：删除失败，8：冻结失败，9：存储中，10：存储失败，11：升级失败
	Status *int32 `json:"status,omitempty"`

	// 标签列表，标签键值map转化的json字符串，如\"{\\\"key\\\":\\\"value\\\"}\"
	Tags *string `json:"tags,omitempty"`
}

func (o FirewallInstanceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirewallInstanceVo struct{}"
	}

	return strings.Join([]string{"FirewallInstanceVo", string(data)}, " ")
}
