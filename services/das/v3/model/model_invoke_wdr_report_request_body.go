package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeWdrReportRequestBody 获取WDR数据请求体
type InvokeWdrReportRequestBody struct {

	// WDR快照开始ID
	StartSnapshotId int64 `json:"start_snapshot_id"`

	// WDR快照结束ID
	EndSnapshotId int64 `json:"end_snapshot_id"`

	// 实例节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`
}

func (o InvokeWdrReportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeWdrReportRequestBody struct{}"
	}

	return strings.Join([]string{"InvokeWdrReportRequestBody", string(data)}, " ")
}
