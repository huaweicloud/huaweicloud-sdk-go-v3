package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTaurusDbNodeProcessesResponse Response Object
type DeleteTaurusDbNodeProcessesResponse struct {

	// **参数解释**：  已终止的用户会话线程ID列表。
	ProcessesKilled *[]int64 `json:"processes_killed,omitempty"`

	// **参数解释**：  不存在的用户会话线程ID列表。
	ProcessesNotFound *[]int64 `json:"processes_not_found,omitempty"`
	HttpStatusCode    int      `json:"-"`
}

func (o DeleteTaurusDbNodeProcessesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaurusDbNodeProcessesResponse struct{}"
	}

	return strings.Join([]string{"DeleteTaurusDbNodeProcessesResponse", string(data)}, " ")
}
