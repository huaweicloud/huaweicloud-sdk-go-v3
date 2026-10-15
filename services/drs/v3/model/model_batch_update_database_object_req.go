package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateDatabaseObjectReq 批量更新数据库对象选择消息体
type BatchUpdateDatabaseObjectReq struct {

	// **参数解释**：  批量更新数据库对象选择的任务信息列表，每个元素包含一个任务的对象选择配置。  **约束限制**：  - 建议批量调用数量不超过10个。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Jobs []UpdateDatabaseObjectReq `json:"jobs"`
}

func (o BatchUpdateDatabaseObjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateDatabaseObjectReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateDatabaseObjectReq", string(data)}, " ")
}
