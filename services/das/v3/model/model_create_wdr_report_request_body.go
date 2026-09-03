package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWdrReportRequestBody 触发WDR请求体
type CreateWdrReportRequestBody struct {

	// 实例节点ID，实例节点的唯一标识。仅支持GaussDB实例节点
	NodeId string `json:"node_id"`

	// WDR快照开始ID
	StartSnapshotId int64 `json:"start_snapshot_id"`

	// WDR快照结束ID
	EndSnapshotId int64 `json:"end_snapshot_id"`

	// 时区。格式：Asia/Shanghai
	TimeZone string `json:"time_zone"`
}

func (o CreateWdrReportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWdrReportRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWdrReportRequestBody", string(data)}, " ")
}
