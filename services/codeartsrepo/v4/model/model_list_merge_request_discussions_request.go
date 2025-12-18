package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestDiscussionsRequest Request Object
type ListMergeRequestDiscussionsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：**  合并请求 iid。
	MergeRequestIid int32 `json:"merge_request_iid"`

	// **参数解释：** 是否筛选解决的检视意见 - true 筛选所有未解决的检视意见 - false 筛选所有已解决的检视意见 - '' 不传此值时默认查询所有检视意见
	Unresolved *string `json:"unresolved,omitempty"`

	// **参数解释：** 操作者id。
	AuthorId *int32 `json:"author_id,omitempty"`

	// **参数解释：** 检视意见返回排序 - asc 按创建时间正序返回 - desc 按创建时间倒序返回
	Sort *string `json:"sort,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMergeRequestDiscussionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestDiscussionsRequest struct{}"
	}

	return strings.Join([]string{"ListMergeRequestDiscussionsRequest", string(data)}, " ")
}
