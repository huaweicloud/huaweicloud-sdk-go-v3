package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenericActionRequestBody 通用Agent行为请求体，通过'action'字段标识具体行为类型和对应的'parameter'内容。
type GenericActionRequestBody struct {

	// 表示具体行为的类型。命名为“t_行为名称”。例如 \"t_dialog_feedback\" 。
	Type string `json:"type"`

	// workspace ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 具体行为的数据内容，其结构由'action'字段决定。 - 回答质量人工反馈action: t_dialog_feedback, 参考DialogFeedbackData结构
	Parameter *interface{} `json:"parameter"`
}

func (o GenericActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenericActionRequestBody struct{}"
	}

	return strings.Join([]string{"GenericActionRequestBody", string(data)}, " ")
}
