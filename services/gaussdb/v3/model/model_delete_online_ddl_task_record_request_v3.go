package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOnlineDdlTaskRecordRequestV3 **参数解释**：  创建无锁变更任务请求体。  **约束限制**：  不涉及。
type DeleteOnlineDdlTaskRecordRequestV3 struct {

	// **参数解释**：   要删除的实例无锁变更任务记录标识。  获取方法请参见[查询无锁变更任务记录列表](https://support.huaweicloud.com/api-taurusdb/ListOnlineDdlTaskRecords.html)。  **约束限制**： 不涉及。   **取值范围**：  不涉及。  **默认取值**： 不涉及。
	TaskId string `json:"task_id"`
}

func (o DeleteOnlineDdlTaskRecordRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOnlineDdlTaskRecordRequestV3 struct{}"
	}

	return strings.Join([]string{"DeleteOnlineDdlTaskRecordRequestV3", string(data)}, " ")
}
