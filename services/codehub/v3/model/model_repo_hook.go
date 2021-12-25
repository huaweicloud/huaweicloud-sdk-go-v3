package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoHook struct {
	// 是否触发build_events事件

	BuildEvents *bool `json:"build_events,omitempty"`
	// 仓库统计创建的时间

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
	// 是否使用ssl验证

	EnableSslVerification *bool `json:"enable_ssl_verification,omitempty"`
	// hook id

	Id *int32 `json:"id,omitempty"`
	// 是否触发issues_events事件

	IssuesEvents *bool `json:"issues_events,omitempty"`
	// 是否触发merge_requests_events事件

	MergeRequestsEvents *bool `json:"merge_requests_events,omitempty"`
	// 是否触发note_events事件

	NoteEvents *bool `json:"note_events,omitempty"`
	// 是否触发pipeline_events事件

	PipelineEvents *bool `json:"pipeline_events,omitempty"`
	// 仓库id

	ProjectId *int32 `json:"project_id,omitempty"`
	// 是否触发push_events事件

	PushEvents *bool `json:"push_events,omitempty"`
	// 是否触发repository_update_events事件

	RepositoryUpdateEvents *bool `json:"repository_update_events,omitempty"`
	// 是否触发tag_push_events事件

	TagPushEvents *bool `json:"tag_push_events,omitempty"`
	// 是否触发wiki_page_events事件

	WikiPageEvents *bool `json:"wiki_page_events,omitempty"`
}

func (o RepoHook) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoHook struct{}"
	}

	return strings.Join([]string{"RepoHook", string(data)}, " ")
}
