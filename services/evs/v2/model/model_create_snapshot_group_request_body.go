package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSnapshotGroupRequestBody 创建快照一致性组的请求体。
type CreateSnapshotGroupRequestBody struct {
	SnapshotGroup *SnapshotGroup `json:"snapshot_group"`
}

func (o CreateSnapshotGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSnapshotGroupRequestBody", string(data)}, " ")
}
