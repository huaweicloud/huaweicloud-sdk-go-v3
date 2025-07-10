package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowTemplateResponse Response Object
type ShowWorkflowTemplateResponse struct {
	Result *WorkflowTemplateVo `json:"result,omitempty"`

	// 状态码
	Status *string `json:"status,omitempty"`

	// 响应信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWorkflowTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowTemplateResponse", string(data)}, " ")
}
