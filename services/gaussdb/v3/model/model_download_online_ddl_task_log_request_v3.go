package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadOnlineDdlTaskLogRequestV3 **参数解释**：  下载实例无锁变更任务日志记录请求体。
type DownloadOnlineDdlTaskLogRequestV3 struct {

	// **参数解释**：   无锁变更任务标识。  获取方法请参见[查询无锁变更任务记录列表](https://support.huaweicloud.com/api-taurusdb/ListOnlineDdlTaskRecords.html)。  **约束限制**： 不涉及。  **取值范围**：  不涉及。   **默认取值**：  不涉及。
	TaskId string `json:"task_id"`
}

func (o DownloadOnlineDdlTaskLogRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadOnlineDdlTaskLogRequestV3 struct{}"
	}

	return strings.Join([]string{"DownloadOnlineDdlTaskLogRequestV3", string(data)}, " ")
}
