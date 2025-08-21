package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleMergeRequestDetailDto struct {

	// **参数解释：** 合并请求id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 合并请求iid。 **取值范围：** 不涉及。
	Iid *int32 `json:"iid,omitempty"`

	// **参数解释：** 标题。 **取值范围：** 不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释：** 描述。 **取值范围：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 源分支。 **取值范围：** 不涉及。
	SourceBranch *string `json:"source_branch,omitempty"`

	// **参数解释：** 目标分支。 **取值范围：** 不涉及。
	TargetBranch *string `json:"target_branch,omitempty"`

	// **参数解释：** 状态。 **取值范围：** 不涉及。
	State *string `json:"state,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 审核模式。 **取值范围：** 不涉及。
	ReviewMode *string `json:"review_mode,omitempty"`

	Author *UserBasicDto `json:"author,omitempty"`

	// **参数解释：** 合并请求类型。 **取值范围：** 不涉及。
	MergeRequestType *string `json:"merge_request_type,omitempty"`

	// **参数解释：** 送审结果。 **取值范围：** - true，成功。 - false，失败
	ModerationResult *bool `json:"moderation_result,omitempty"`

	// **参数解释：** 送审时间时间戳。 **取值范围：** 不涉及。
	ModerationTime *int64 `json:"moderation_time,omitempty"`

	// **参数解释：** 送审状态码。 **取值范围：** 不涉及。
	ModerationStatus *int32 `json:"moderation_status,omitempty"`

	// **参数解释：** 是否使用临时分支。 **取值范围：** - true，使用。 - false，不使用
	IsUseTempBranch *bool `json:"is_use_temp_branch,omitempty"`
}

func (o SimpleMergeRequestDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleMergeRequestDetailDto struct{}"
	}

	return strings.Join([]string{"SimpleMergeRequestDetailDto", string(data)}, " ")
}
