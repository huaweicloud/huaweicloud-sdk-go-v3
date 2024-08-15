package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceCompliant 节点合规性报告
type InstanceCompliant struct {
	CompliantSummary *CompliantSummary `json:"compliant_summary,omitempty"`

	NonCompliantSummary *NonCompliantSummary `json:"non_compliant_summary,omitempty"`

	ExecutionSummary *ExecutionSummary `json:"execution_summary,omitempty"`

	// id
	Id *string `json:"id,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 节点名称
	Name *string `json:"name,omitempty"`

	// 节点ID
	InstanceId *string `json:"instance_id,omitempty"`

	// cce集群节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 节点IP
	Ip *string `json:"ip,omitempty"`

	// 弹性公网ip
	Eip *string `json:"eip,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 分组
	Group *string `json:"group,omitempty"`

	// 报告场景(CCE,ECS)
	ReportScene *string `json:"report_scene,omitempty"`

	// cce 集群信息id
	CceInfoId *string `json:"cce_info_id,omitempty"`

	// 合规性状态
	Status *string `json:"status,omitempty"`

	// 基线id
	BaselineId *string `json:"baseline_id,omitempty"`

	// 基线名称
	BaselineName *string `json:"baseline_name,omitempty"`

	// 基线规则类型
	RuleType *string `json:"rule_type,omitempty"`

	// 操作系统
	OperatingSystem *string `json:"operating_system,omitempty"`
}

func (o InstanceCompliant) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceCompliant struct{}"
	}

	return strings.Join([]string{"InstanceCompliant", string(data)}, " ")
}
