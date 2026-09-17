package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SnapshotTaskStatus struct {

	// **参数解释：** 最近一次备份的时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	LatestBackupTime *string `json:"latestBackupTime,omitempty"`
}

func (o SnapshotTaskStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotTaskStatus struct{}"
	}

	return strings.Join([]string{"SnapshotTaskStatus", string(data)}, " ")
}
