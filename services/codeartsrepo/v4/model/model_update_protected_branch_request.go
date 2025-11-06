package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProtectedBranchRequest Request Object
type UpdateProtectedBranchRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 保护分支名或通配符表达式。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	BranchName string `json:"branch_name"`

	// 保护分支权限列表
	Body *[]ProtectedBranchProtectedActionBodyDto `json:"body,omitempty"`
}

func (o UpdateProtectedBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectedBranchRequest struct{}"
	}

	return strings.Join([]string{"UpdateProtectedBranchRequest", string(data)}, " ")
}
