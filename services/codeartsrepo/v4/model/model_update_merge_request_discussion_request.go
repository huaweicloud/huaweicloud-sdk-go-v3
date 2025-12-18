package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMergeRequestDiscussionRequest Request Object
type UpdateMergeRequestDiscussionRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：**  合并请求 iid。
	MergeRequestIid int32 `json:"merge_request_iid"`

	// **参数解释：** 检视意见id(主评和回复共用)。 **取值范围：** 字符串长度40。
	DiscussionId string `json:"discussion_id"`

	// **参数解释：** 评论id(主评论和回复不共用)。
	NoteId int32 `json:"note_id"`

	Body *UpdateMergeRequestNoteDto `json:"body,omitempty"`
}

func (o UpdateMergeRequestDiscussionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMergeRequestDiscussionRequest struct{}"
	}

	return strings.Join([]string{"UpdateMergeRequestDiscussionRequest", string(data)}, " ")
}
