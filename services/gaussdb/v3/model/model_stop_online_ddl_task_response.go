package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopOnlineDdlTaskResponse Response Object
type StopOnlineDdlTaskResponse struct {

	// **参数解释**：  工作流ID，停止实例无锁变更任务的工作流标识。  **取值范围**：   不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopOnlineDdlTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopOnlineDdlTaskResponse struct{}"
	}

	return strings.Join([]string{"StopOnlineDdlTaskResponse", string(data)}, " ")
}
