package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbMaskTaskRequest Request Object
type ListDbMaskTaskRequest struct {

	// 模板ID
	TemplateId string `json:"template_id"`

	// 工作区ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDbMaskTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbMaskTaskRequest struct{}"
	}

	return strings.Join([]string{"ListDbMaskTaskRequest", string(data)}, " ")
}
