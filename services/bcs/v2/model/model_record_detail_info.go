package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作记录详细信息
type RecordDetailInfo struct {
	// 操作记录ID

	OperationId *string `json:"operation_id,omitempty"`
	// 资源类型

	ResourceType *string `json:"resource_type,omitempty"`
	// 操作类型

	OperationType *string `json:"operation_type,omitempty"`
	// 租户ID

	DomainId *string `json:"domain_id,omitempty"`
	// 项目ID

	ProjectId *string `json:"project_id,omitempty"`
	// 区块链ID

	BlockchainId *string `json:"blockchain_id,omitempty"`
	// 区块链名称

	BlockchainName *string `json:"blockchain_name,omitempty"`

	ClusterInfo *OprecordCluster `json:"cluster_info,omitempty"`
	// 操作流程，key为流程名，value为流程信息

	OperationProcess map[string]ProcessInfo `json:"operation_process,omitempty"`
	// 记录更新时间

	RecordTime *int64 `json:"record_time,omitempty"`
	// 操作状态

	OperationStatus *string `json:"operation_status,omitempty"`
	// 操作过程信息记录

	Message *[]string `json:"message,omitempty"`
	// 操作描述

	Desc *string `json:"desc,omitempty"`
}

func (o RecordDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordDetailInfo struct{}"
	}

	return strings.Join([]string{"RecordDetailInfo", string(data)}, " ")
}
