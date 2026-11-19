package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRetainedBackupsResponse Response Object
type ListRetainedBackupsResponse struct {

	// **参数解释**：  保留备份信息列表。
	RecycleBackupInfos *[]RetainedBackupInfo `json:"recycle_backup_infos,omitempty"`

	Policy *RetainedBackupPolicy `json:"policy,omitempty"`

	// **参数解释**：  保留备份总数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRetainedBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRetainedBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListRetainedBackupsResponse", string(data)}, " ")
}
