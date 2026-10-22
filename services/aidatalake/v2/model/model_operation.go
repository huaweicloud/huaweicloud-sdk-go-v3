package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Operation **参数解释**：资源操作记录，包含操作类型、操作结果。
type Operation struct {

	// **参数解释**：操作类型。 **取值范围**：   - CREATE：创建。   - SCALEIN：缩容。   - SCALEOUT：扩容。   - DELETE：删除。
	OperationType *string `json:"operation_type,omitempty"`

	// **参数解释**：操作结果。 **取值范围**：   - EXECUTING：执行中。   - SUCCESS：操作成功。   - FAILED：操作失败。
	OperationResult *string `json:"operation_result,omitempty"`
}

func (o Operation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Operation struct{}"
	}

	return strings.Join([]string{"Operation", string(data)}, " ")
}
