package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateRepositoryUserGroupRequest Request Object
type AssociateRepositoryUserGroupRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 仓库id
	RepositoryId int32 `json:"repository_id"`

	// 成员组id
	UserGroupId string `json:"user_group_id"`
}

func (o AssociateRepositoryUserGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRepositoryUserGroupRequest struct{}"
	}

	return strings.Join([]string{"AssociateRepositoryUserGroupRequest", string(data)}, " ")
}
