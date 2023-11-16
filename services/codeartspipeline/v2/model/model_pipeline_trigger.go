package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineTrigger struct {

	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty"`

	// git链接
	GitUrl *string `json:"git_url,omitempty"`

	// 代码仓类型
	GitType *string `json:"git_type,omitempty"`

	// 是否自动提交
	IsAutoCommit *bool `json:"is_auto_commit,omitempty"`

	// 事件
	Events *[]CodeEvent `json:"events,omitempty"`

	// 回调id
	HookId *string `json:"hook_id,omitempty"`

	// 仓库id
	RepoId *string `json:"repo_id,omitempty"`

	// 扩展点id
	EndpointId *string `json:"endpoint_id,omitempty"`

	// 回调链接
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 用户token
	SecurityToken *string `json:"security_token,omitempty"`
}

func (o PipelineTrigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineTrigger struct{}"
	}

	return strings.Join([]string{"PipelineTrigger", string(data)}, " ")
}
