package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineRunsPageBuildParams **参数解释**： 构建参数。 **取值范围**： 不涉及。
type ListPipelineRunsPageBuildParams struct {

	// **参数解释**： 合并请求事件类型。 **取值范围**： - open：打开。 - reopen：重开。 - update：更新。 - merge：合并。
	Action *string `json:"action,omitempty"`

	// **参数解释**： 基于分支还是标签触发。 **取值范围**： - branch：分支触发。 - tag：标签触发。
	BuildType *string `json:"build_type,omitempty"`

	// **参数解释**： 代码仓提交ID。 **取值范围**： 不涉及。
	CommitId *string `json:"commit_id,omitempty"`

	// **参数解释**： 运行事件类型。 **取值范围**： - Manual：手动触发。 - Scheduler：定时任务。 - MR：MR触发。 - Push：Push事件触发。 - CreateTag：Tag事件触发。 - Issue：Issue触发。 - Note：评论触发。
	EventType *string `json:"event_type,omitempty"`

	// **参数解释**： 合并请求ID。 **取值范围**： 不涉及。
	MergeId *string `json:"merge_id,omitempty"`

	// **参数解释**： 代码仓提交信息。 **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释**： 源分支。 **取值范围**： 不涉及。
	SourceBranch *string `json:"source_branch,omitempty"`

	// **参数解释**： 标签。 **取值范围**： 不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**： 目标分支。 **取值范围**： 不涉及。
	TargetBranch *string `json:"target_branch,omitempty"`

	// **参数解释**： Repo代码仓ID。 **取值范围**： 不涉及。
	CodehubId *string `json:"codehub_id,omitempty"`

	// **参数解释**： 代码仓https地址。 **取值范围**： 不涉及。
	GitUrl *string `json:"git_url,omitempty"`

	// **参数解释**： 源Repo代码仓ID。 **取值范围**： 不涉及。
	SourceCodehubId *string `json:"source_codehub_id,omitempty"`

	// **参数解释**： 源Repo代码仓地址。 **取值范围**： 不涉及。
	SourceCodehubUrl *string `json:"source_codehub_url,omitempty"`

	// **参数解释**： 源Repo代码仓http地址。 **取值范围**： 不涉及。
	SourceCodehubHttpUrl *string `json:"source_codehub_http_url,omitempty"`
}

func (o ListPipelineRunsPageBuildParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineRunsPageBuildParams struct{}"
	}

	return strings.Join([]string{"ListPipelineRunsPageBuildParams", string(data)}, " ")
}
