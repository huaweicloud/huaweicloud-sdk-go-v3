package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOnlineDdlTaskRecordsResponse Response Object
type ListOnlineDdlTaskRecordsResponse struct {

	// **参数解释**：  无锁变更任务详情列表。
	Records *[]RecordItem `json:"records,omitempty"`

	// **参数解释**：   无锁变更任务记录总数，整数。  **取值范围**：   ≥0。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOnlineDdlTaskRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOnlineDdlTaskRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListOnlineDdlTaskRecordsResponse", string(data)}, " ")
}
