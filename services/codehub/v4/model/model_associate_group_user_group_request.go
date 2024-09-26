package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateGroupUserGroupRequest Request Object
type AssociateGroupUserGroupRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 代码组 ID
	GroupId int32 `json:"group_id"`

	// 成员组 ID
	UserGroupId string `json:"user_group_id"`
}

func (o AssociateGroupUserGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateGroupUserGroupRequest struct{}"
	}

	return strings.Join([]string{"AssociateGroupUserGroupRequest", string(data)}, " ")
}
