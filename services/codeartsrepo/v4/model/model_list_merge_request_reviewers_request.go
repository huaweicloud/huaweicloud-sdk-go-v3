package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestReviewersRequest Request Object
type ListMergeRequestReviewersRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 目标分支。创建MR时，代码将要合入的分支。
	TargetBranch *string `json:"target_branch,omitempty"`

	// **参数解释：** 目标分支。创建MR时，变更代码所属的分支。
	SourceBranch *string `json:"source_branch,omitempty"`

	// **参数解释：**  合并请求 iid。
	MergeRequestIid *int32 `json:"merge_request_iid,omitempty"`

	// **参数解释：** 目标仓库id。创建MR时，代码将要合入的仓库。
	TargetRepositoryId *string `json:"target_repository_id,omitempty"`

	// **参数解释：** 查询关键字，可模糊匹配用户名称、用户昵称、租户名称。
	Search *string `json:"search,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMergeRequestReviewersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestReviewersRequest struct{}"
	}

	return strings.Join([]string{"ListMergeRequestReviewersRequest", string(data)}, " ")
}
