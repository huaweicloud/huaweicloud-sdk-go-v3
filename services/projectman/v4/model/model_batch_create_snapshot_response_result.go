package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSnapshotResponseResult 批量创建快照的结果。
type BatchCreateSnapshotResponseResult struct {

	// 创建成功的快照列表。
	Success *[]CreateSnapshotResult `json:"success,omitempty"`

	// 创建失败的快照列表。
	Failed *[]CreateSnapshotResult `json:"failed,omitempty"`
}

func (o BatchCreateSnapshotResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSnapshotResponseResult struct{}"
	}

	return strings.Join([]string{"BatchCreateSnapshotResponseResult", string(data)}, " ")
}
