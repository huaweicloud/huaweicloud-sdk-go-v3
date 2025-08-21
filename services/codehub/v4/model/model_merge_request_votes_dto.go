package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MergeRequestVotesDto 合并请求打分记录(单人)
type MergeRequestVotesDto struct {

	// **参数解释：** 打分记录的id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 分数。
	Score *int32 `json:"score,omitempty"`

	// **参数解释：** 作者名。
	AuthorName *string `json:"author_name,omitempty"`

	// **参数解释：** 作者用户名。
	AuthorUsername *string `json:"author_username,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 最后一次提交记录的id。
	LastCommittedId *string `json:"last_committed_id,omitempty"`

	// **参数解释：** 作者id。
	AuthorId *int32 `json:"author_id,omitempty"`

	// **参数解释：** 作者头像url。
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// **参数解释：** 作者昵称。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 作者租户名称。
	TenantName *string `json:"tenant_name,omitempty"`
}

func (o MergeRequestVotesDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestVotesDto struct{}"
	}

	return strings.Join([]string{"MergeRequestVotesDto", string(data)}, " ")
}
