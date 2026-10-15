package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobResp 创建任务响应体
type CreateJobResp struct {

	// **参数解释**：  任务ID，用于唯一标识一个数据复制任务，在查询、修改、删除等操作中作为任务的识别凭证。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	Id string `json:"id"`

	// **参数解释**：  任务名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：  任务当前运行状态。  **约束限制**：  不涉及。  **取值范围**：  - CREATING：创建中 - CREATE_FAILED：创建失败 - CONFIGURATION：配置中 - WAITING_FOR_START：等待启动中 - RELEASE_RESOURCE_COMPLETE：已结束 - DELETED：已删除 - INCRE_TRANSFER_STARTED：增量迁移中 - INCRE_TRANSFER_FAILED：增量迁移失败 - FULL_TRANSFER_STARTED：全量迁移中 - FULL_TRANSFER_COMPLETE：全量迁移完成 - PAUSING：暂停中 - FULL_TRANSFER_FAILED：全量迁移失败  **默认取值**：  不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**：  任务的创建时间，时间戳格式。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**：  错误码，当任务报错时返回的错误代码。  **约束限制**：  不涉及。  **取值范围**：  格式为DRS.XXXXXX。  **默认取值**：  不涉及
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误详细信息，当任务执行失败时返回的具体错误描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`

	// **参数解释**：  子任务ID列表，有子任务时返回该字段。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ChildIds *[]string `json:"child_ids,omitempty"`
}

func (o CreateJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobResp struct{}"
	}

	return strings.Join([]string{"CreateJobResp", string(data)}, " ")
}
