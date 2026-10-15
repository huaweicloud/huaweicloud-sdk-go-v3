package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseObjectResp 对象选择返回体
type DatabaseObjectResp struct {

	// **参数解释**：  迁移任务ID，用于唯一标识一个数据复制任务，在查询、修改、删除等操作中作为任务的识别凭证。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：  对象选择任务是否成功的标志。  **约束限制**：  不涉及。  **取值范围**：  - true：对象选择成功 - false：对象选择失败  **默认取值**：  不涉及
	Status *bool `json:"status,omitempty"`

	// **参数解释**：  错误码，当任务报错时返回的错误代码。  **约束限制**：  不涉及。  **取值范围**：  格式为DRS.XXXXXX。  **默认取值**：  不涉及
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误详细信息，当任务执行失败时返回的具体错误描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o DatabaseObjectResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseObjectResp struct{}"
	}

	return strings.Join([]string{"DatabaseObjectResp", string(data)}, " ")
}
