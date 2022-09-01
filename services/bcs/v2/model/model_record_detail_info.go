package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作记录详细信息
type RecordDetailInfo struct {

	// 操作记录ID
	OperationId *string `json:"operation_id,omitempty" xml:"operation_id"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type"`

	// 操作类型
	OperationType *string `json:"operation_type,omitempty" xml:"operation_type"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 区块链ID
	BlockchainId *string `json:"blockchain_id,omitempty" xml:"blockchain_id"`

	// 区块链名称
	BlockchainName *string `json:"blockchain_name,omitempty" xml:"blockchain_name"`

	ClusterInfo *OprecordCluster `json:"cluster_info,omitempty" xml:"cluster_info"`

	// 操作流程，key为流程名，value为流程信息
	OperationProcess map[string]ProcessInfo `json:"operation_process,omitempty" xml:"operation_process"`

	// 记录更新时间
	RecordTime *int64 `json:"record_time,omitempty" xml:"record_time"`

	// 操作状态
	OperationStatus *string `json:"operation_status,omitempty" xml:"operation_status"`

	// 操作过程信息记录
	Message *[]string `json:"message,omitempty" xml:"message"`

	// 操作描述
	Desc *string `json:"desc,omitempty" xml:"desc"`
}

func (o RecordDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordDetailInfo struct{}"
	}

	return strings.Join([]string{"RecordDetailInfo", string(data)}, " ")
}
