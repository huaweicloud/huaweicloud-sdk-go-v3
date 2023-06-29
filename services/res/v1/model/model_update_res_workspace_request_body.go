package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResWorkspaceRequestBody This is a auto create Body Object
type UpdateResWorkspaceRequestBody struct {

	// 工作空间名称，1-64位数字、字母、下划线、中划线组成，支持中文，不能以 - 结尾。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`
}

func (o UpdateResWorkspaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResWorkspaceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResWorkspaceRequestBody", string(data)}, " ")
}
