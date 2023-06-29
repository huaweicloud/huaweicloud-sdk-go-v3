package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResWorkspaceRequestBody This is a auto create Body Object
type CreateResWorkspaceRequestBody struct {

	// 工作空间名称，1-64位的数字、字母、下划线、中划线组成，支持中文，不能以 - 结尾。
	Name string `json:"name"`

	// 企业项目编号。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 描述。
	Description *string `json:"description,omitempty"`
}

func (o CreateResWorkspaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResWorkspaceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResWorkspaceRequestBody", string(data)}, " ")
}
