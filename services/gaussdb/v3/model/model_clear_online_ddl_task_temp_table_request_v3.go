package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearOnlineDdlTaskTempTableRequestV3 **参数解释**：  清理实例无锁变更任务临时表请求体。
type ClearOnlineDdlTaskTempTableRequestV3 struct {

	// **参数解释**：   无锁变更任务唯一标识。  获取方法参见[查询无锁变更任务记录列表](https://support.huaweicloud.com/api-taurusdb/ListOnlineDdlTaskRecords.html)。   **约束限制**：   不涉及。   **取值范围**：   不涉及。  **默认取值**：   不涉及。
	TaskId string `json:"task_id"`

	// **参数解释**：  无锁变更任务详细内容，包含目标数据库和临时表名。  **约束限制**：  不涉及。
	TaskContent []TaskContentItem `json:"task_content"`
}

func (o ClearOnlineDdlTaskTempTableRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearOnlineDdlTaskTempTableRequestV3 struct{}"
	}

	return strings.Join([]string{"ClearOnlineDdlTaskTempTableRequestV3", string(data)}, " ")
}
