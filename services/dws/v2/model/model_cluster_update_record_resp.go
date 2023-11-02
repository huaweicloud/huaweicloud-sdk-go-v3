package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterUpdateRecordResp 集群升级记录响应体
type ClusterUpdateRecordResp struct {

	// 升级项目ID
	ItemId *string `json:"item_id,omitempty"`

	// 升级状态
	Status *string `json:"status,omitempty"`

	// 升级类型
	RecordType *string `json:"record_type,omitempty"`

	// 升级前版本
	FromVersion *string `json:"from_version,omitempty"`

	// 目标版本
	ToVersion *string `json:"to_version,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 升级任务ID
	JobId *string `json:"job_id,omitempty"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o ClusterUpdateRecordResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterUpdateRecordResp struct{}"
	}

	return strings.Join([]string{"ClusterUpdateRecordResp", string(data)}, " ")
}
