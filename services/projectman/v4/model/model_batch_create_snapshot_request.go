package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSnapshotRequest 批量创建工作项快照请求对象
type BatchCreateSnapshotRequest struct {

	// 需要创建快照的工作项数组。 每次最多支持对50个工作项创建快照。
	Issues []BatchCreateSnapshotRequestIssues `json:"issues"`
}

func (o BatchCreateSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSnapshotRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSnapshotRequest", string(data)}, " ")
}
