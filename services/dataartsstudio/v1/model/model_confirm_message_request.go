package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmMessageRequest Request Object
type ConfirmMessageRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	Body *OpenApiParaForCheckMessage `json:"body,omitempty"`
}

func (o ConfirmMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmMessageRequest struct{}"
	}

	return strings.Join([]string{"ConfirmMessageRequest", string(data)}, " ")
}
