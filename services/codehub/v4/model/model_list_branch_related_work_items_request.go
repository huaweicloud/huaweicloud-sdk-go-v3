package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBranchRelatedWorkItemsRequest Request Object
type ListBranchRelatedWorkItemsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 分支名称。  **约束限制：** 不支持以 - . refs/heads/ refs/remotes/ 开头，不支持空格 [ \\ < ~ ^ : ? * ! ( ) ' \" | 等特殊字符，不支持以. / .lock结尾。  **取值范围：** 字符串长度不少于1，不超过200。 **默认取值：** 不涉及。
	BranchName string `json:"branch_name"`
}

func (o ListBranchRelatedWorkItemsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBranchRelatedWorkItemsRequest struct{}"
	}

	return strings.Join([]string{"ListBranchRelatedWorkItemsRequest", string(data)}, " ")
}
