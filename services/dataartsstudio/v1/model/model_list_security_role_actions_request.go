package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityRoleActionsRequest Request Object
type ListSecurityRoleActionsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *RoleActionQueryDto `json:"body,omitempty"`
}

func (o ListSecurityRoleActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityRoleActionsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityRoleActionsRequest", string(data)}, " ")
}
