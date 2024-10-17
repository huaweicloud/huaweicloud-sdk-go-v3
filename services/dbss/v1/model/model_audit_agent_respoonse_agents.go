package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditAgentRespoonseAgents struct {

	// agent ID
	AgentId string `json:"agent_id"`

	// agent 类型
	AgentType string `json:"agent_type"`

	// agent OS
	AgentOs string `json:"agent_os"`

	// agent安装节点IP
	AgentIp string `json:"agent_ip"`

	// 内存阈值
	MemThreshold *int32 `json:"mem_threshold,omitempty"`

	// cpu阈值
	CpuThreshold *int32 `json:"cpu_threshold,omitempty"`

	// agent状态
	Status *int32 `json:"status,omitempty"`

	// agent网卡
	AgentNic *string `json:"agent_nic,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 数据流量抓取状态
	DatacapStatus *int32 `json:"datacap_status,omitempty"`

	// agent安装地址
	AgentUrl *string `json:"agent_url,omitempty"`

	// 是否CCE场景
	Universal *bool `json:"universal,omitempty"`

	// sha256值
	Sha256 *string `json:"sha256,omitempty"`
}

func (o AuditAgentRespoonseAgents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditAgentRespoonseAgents struct{}"
	}

	return strings.Join([]string{"AuditAgentRespoonseAgents", string(data)}, " ")
}
