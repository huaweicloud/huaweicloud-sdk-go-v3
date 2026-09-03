package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRetainBackupsResponse Response Object
type ListRetainBackupsResponse struct {
	Backups *RetainBackup `json:"backups,omitempty"`

	// **参数解释**：  保留备份总数  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRetainBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRetainBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListRetainBackupsResponse", string(data)}, " ")
}
