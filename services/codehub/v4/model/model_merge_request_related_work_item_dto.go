package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestRelatedWorkItemDto struct {

	// 关联id
	Id *int32 `json:"id,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 仓库id
	RepoId *int32 `json:"repo_id,omitempty"`

	// 合并请求id
	MergeRequestId *int32 `json:"merge_request_id,omitempty"`

	// 目标分支
	TargetBranch *string `json:"target_branch,omitempty"`

	// 源分支
	SourceBranch *string `json:"source_branch,omitempty"`

	// 合并请求标题
	MergeRequestTitle *string `json:"merge_request_title,omitempty"`

	// 合并请求url
	MergeRequestUrl *string `json:"merge_request_url,omitempty"`

	// 合并请求状态
	MergeRequestState *string `json:"merge_request_state,omitempty"`

	// 关联单id
	RelatedId *string `json:"related_id,omitempty"`

	// 关联单标题
	RelatedTitle *string `json:"related_title,omitempty"`

	// 关联单url
	RelatedUrl *string `json:"related_url,omitempty"`

	// 关联结果
	Result *int32 `json:"result,omitempty"`

	// 关联结果信息
	Message *string `json:"message,omitempty"`

	// 创建时间
	CreateAt *string `json:"create_at,omitempty"`

	// 更新时间
	UpdateAt *string `json:"update_at,omitempty"`

	// 话题
	Subject *string `json:"subject,omitempty"`

	// 跟单人
	Tracker *interface{} `json:"tracker,omitempty"`

	// 状态
	Status *interface{} `json:"status,omitempty"`
}

func (o MergeRequestRelatedWorkItemDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestRelatedWorkItemDto struct{}"
	}

	return strings.Join([]string{"MergeRequestRelatedWorkItemDto", string(data)}, " ")
}
