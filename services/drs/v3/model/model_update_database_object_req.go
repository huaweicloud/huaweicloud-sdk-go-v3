package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatabaseObjectReq 更新对象选择信息体
type UpdateDatabaseObjectReq struct {

	// **参数解释**：  迁移任务ID，用于唯一标识一个数据复制任务，在查询、修改、删除等操作中作为任务的识别凭证。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	JobId string `json:"job_id"`

	// **参数解释**：  是否进行对象选择，决定是自定义选择迁移对象还是迁移全部对象。  **约束限制**：  不涉及。  **取值范围**：  - true：自定义迁移对象，需要通过job字段指定具体迁移对象 - false：全部迁移，迁移源库所有对象  **默认取值**：  false
	Selected *bool `json:"selected,omitempty"`

	// **参数解释**：  是否库级同步，指定同步任务是否以库为粒度进行对象选择。  **约束限制**：  不涉及。  **取值范围**：  - true：库级同步 - false：非库级同步  **默认取值**：  不涉及
	SyncDatabase *bool `json:"sync_database,omitempty"`

	// **参数解释**：  数据对象选择信息列表，包含需要迁移的数据库、表、视图等对象的详细配置。  **约束限制**：  - selected为true时必填。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Job *[]DatabaseInfo `json:"job,omitempty"`
}

func (o UpdateDatabaseObjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseObjectReq struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseObjectReq", string(data)}, " ")
}
