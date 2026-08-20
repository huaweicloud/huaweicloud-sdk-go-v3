package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FailureDetail 批量操作快照失败响应对象。
type FailureDetail struct {

	// 快照ID。
	SnapshotId *string `json:"snapshot_id,omitempty"`

	// 快照标题。
	SnapshotTitle *string `json:"snapshot_title,omitempty"`

	// 失败原因。
	FailureReason *string `json:"failure_reason,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o FailureDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailureDetail struct{}"
	}

	return strings.Join([]string{"FailureDetail", string(data)}, " ")
}
