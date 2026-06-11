package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupSpaceUsage struct {

	// **参数解释：** 备份空间使用量。 **取值范围：** 不涉及。
	BackupUsage *int64 `json:"backup_usage,omitempty"`
}

func (o BackupSpaceUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupSpaceUsage struct{}"
	}

	return strings.Join([]string{"BackupSpaceUsage", string(data)}, " ")
}
