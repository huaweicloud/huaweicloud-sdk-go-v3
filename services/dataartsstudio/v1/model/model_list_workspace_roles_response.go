package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspaceRolesResponse Response Object
type ListWorkspaceRolesResponse struct {

	// 获取DataArtsStudio工作空间角色列表信息
	Body           *[]ApigRoleVo `json:"body,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListWorkspaceRolesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspaceRolesResponse struct{}"
	}

	return strings.Join([]string{"ListWorkspaceRolesResponse", string(data)}, " ")
}
