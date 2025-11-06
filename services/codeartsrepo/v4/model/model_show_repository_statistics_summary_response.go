package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryStatisticsSummaryResponse Response Object
type ShowRepositoryStatisticsSummaryResponse struct {

	// **参数解释：** 分支数量。 **取值范围：** 不涉及。
	BranchesCount *int32 `json:"branches_count,omitempty"`

	// **参数解释：** 提交数量。 **取值范围：** 不涉及。
	CommitsCount *int32 `json:"commits_count,omitempty"`

	// **参数解释：** 成员数量。
	MembersCount *int32 `json:"members_count,omitempty"`

	// **参数解释：** Tag数量。 **取值范围：** 不涉及。
	TagsCount *int32 `json:"tags_count,omitempty"`

	// **参数解释：** 合并请求数量。 **取值范围：** 不涉及。
	MergeRequestCount *int64 `json:"merge_request_count,omitempty"`

	// **参数解释：** 备注数量。 **取值范围：** 不涉及。
	NoteCount      *int64 `json:"note_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRepositoryStatisticsSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryStatisticsSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryStatisticsSummaryResponse", string(data)}, " ")
}
