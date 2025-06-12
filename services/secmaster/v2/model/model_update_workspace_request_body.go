package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkspaceRequestBody 更新工作空间时参数对象
type UpdateWorkspaceRequestBody struct {

	// 工作空间名称
	Name *string `json:"name,omitempty"`

	// 工作空间描述
	Description *string `json:"description,omitempty"`

	// 视图绑定的空间id
	ViewBindId *string `json:"view_bind_id,omitempty"`
}

func (o UpdateWorkspaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkspaceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateWorkspaceRequestBody", string(data)}, " ")
}
