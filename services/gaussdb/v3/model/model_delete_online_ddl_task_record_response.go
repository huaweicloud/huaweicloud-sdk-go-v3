package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOnlineDdlTaskRecordResponse Response Object
type DeleteOnlineDdlTaskRecordResponse struct {

	// **参数解释**：  工作流ID，删除无锁变更任务的工作流标识。   **取值范围**：   不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteOnlineDdlTaskRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOnlineDdlTaskRecordResponse struct{}"
	}

	return strings.Join([]string{"DeleteOnlineDdlTaskRecordResponse", string(data)}, " ")
}
