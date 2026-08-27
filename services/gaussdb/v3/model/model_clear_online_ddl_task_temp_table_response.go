package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearOnlineDdlTaskTempTableResponse Response Object
type ClearOnlineDdlTaskTempTableResponse struct {

	// **参数解释**：  任务流ID，清理实例无锁变更任务临时表的工作流标识。   **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ClearOnlineDdlTaskTempTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearOnlineDdlTaskTempTableResponse struct{}"
	}

	return strings.Join([]string{"ClearOnlineDdlTaskTempTableResponse", string(data)}, " ")
}
