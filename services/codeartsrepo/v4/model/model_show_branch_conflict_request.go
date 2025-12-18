package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBranchConflictRequest Request Object
type ShowBranchConflictRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 查询指定源仓库的数据。
	SourceRepositoryId int32 `json:"source_repository_id"`

	// **参数解释：** 源分支。
	SourceBranch string `json:"source_branch"`

	// **参数解释：** 目标分支。
	TargetBranch string `json:"target_branch"`

	// **参数解释：** 目标仓库id。创建MR时，代码将要合入的仓库。
	TargetRepositoryId int32 `json:"target_repository_id"`
}

func (o ShowBranchConflictRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBranchConflictRequest struct{}"
	}

	return strings.Join([]string{"ShowBranchConflictRequest", string(data)}, " ")
}
