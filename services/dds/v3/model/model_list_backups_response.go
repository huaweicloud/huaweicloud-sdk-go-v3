package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupsResponse Response Object
type ListBackupsResponse struct {

	// **参数解释：** 备份列表。 **取值范围：** 不涉及。
	Backups *[]BackupForList `json:"backups,omitempty"`

	// **参数解释：** 总记录数。 **取值范围：** 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupsResponse", string(data)}, " ")
}
