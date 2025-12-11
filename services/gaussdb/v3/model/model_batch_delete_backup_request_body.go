package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBackupRequestBody 批量删除手动备份请求体
type BatchDeleteBackupRequestBody struct {

	// **参数解释**：  需要删除的手动备份ID列表。 通过调用备份管理[查询全量备份列表](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlBackupList.html)接口获取。  **约束限制**：   列表数量小于等于50。  **取值范围**：  不涉及。   **默认取值**：   不涉及。
	BackupIds []string `json:"backup_ids"`
}

func (o BatchDeleteBackupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBackupRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteBackupRequestBody", string(data)}, " ")
}
