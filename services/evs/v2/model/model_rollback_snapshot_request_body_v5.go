package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackSnapshotRequestBodyV5 回滚快照请求体。
type RollbackSnapshotRequestBodyV5 struct {
	Rollback *RollbackSnapshotRequestBodyNew `json:"rollback"`
}

func (o RollbackSnapshotRequestBodyV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackSnapshotRequestBodyV5 struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotRequestBodyV5", string(data)}, " ")
}
