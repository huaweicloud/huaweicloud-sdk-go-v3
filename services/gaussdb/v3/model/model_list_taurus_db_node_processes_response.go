package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaurusDbNodeProcessesResponse Response Object
type ListTaurusDbNodeProcessesResponse struct {

	// **参数解释**：  用户会话线程信息列表。
	Processes *[]TaurusDbProcessInfo `json:"processes,omitempty"`

	// **参数解释**：  节点中的用户会话线程总数。  **取值范围**：  ≥0
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTaurusDbNodeProcessesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaurusDbNodeProcessesResponse struct{}"
	}

	return strings.Join([]string{"ListTaurusDbNodeProcessesResponse", string(data)}, " ")
}
