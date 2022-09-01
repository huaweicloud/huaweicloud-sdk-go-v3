package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type RollbackSnapshotRequestBody struct {
	Rollback *RollbackSnapshotOption `json:"rollback" xml:"rollback"`
}

func (o RollbackSnapshotRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotRequestBody", string(data)}, " ")
}
