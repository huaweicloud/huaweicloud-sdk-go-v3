package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionBaseResp 操作任务响应体。
type ActionBaseResp struct {

	// **参数解释**：  错误码，当任务报错时返回的错误代码。  **约束限制**：  - 最小长度：12 - 最大长度：12  **取值范围**：  格式为DRS.XXXXXX。  **默认取值**：  不涉及
	ErrorCode string `json:"error_code"`

	// **参数解释**：  错误详细信息，当任务执行失败时返回的具体错误描述。  **约束限制**：  - 最小长度：1 - 最大长度：512  **取值范围**：  不涉及。  **默认取值**：  不涉及
	ErrorMsg string `json:"error_msg"`

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 操作结果。 success，failed
	Status string `json:"status"`
}

func (o ActionBaseResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionBaseResp struct{}"
	}

	return strings.Join([]string{"ActionBaseResp", string(data)}, " ")
}
