package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineTrigger struct {

	// **参数解释**： 流水线ID，可以通过[查询流水线列表](ListPipelines.xml)接口，其中pipelines.pipelineId即为流水线ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	PipelineId *string `json:"pipeline_id,omitempty"`

	// **参数解释**： 代码仓git链接。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GitUrl *string `json:"git_url,omitempty"`

	// **参数解释**： git代码仓类型。 **约束限制**： 不涉及。 **取值范围**： - codehub。 - gitee。 - github。 - gitcode。 - gitlab。 **默认取值**： 不涉及。
	GitType *string `json:"git_type,omitempty"`

	// **参数解释**： 是否自动提交。gitee仓库特有，webhook触发流水线后，自动添加评论。 **约束限制**： 不涉及。 **取值范围**： - true：是自动提交。 - false：不是自动提交。 **默认取值**： 不涉及。
	IsAutoCommit *bool `json:"is_auto_commit,omitempty"`

	// **参数解释**： 触发事件列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Events *[]CodeEvent `json:"events,omitempty"`

	// **参数解释**： 系统生成的回调ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	HookId *string `json:"hook_id,omitempty"`

	// **参数解释**： Repo仓库ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RepoId *string `json:"repo_id,omitempty"`

	// **参数解释**： 代码源扩展点ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**： 回调链接，注册Webhook时生成。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CallbackUrl *string `json:"callback_url,omitempty"`

	// **参数解释**： 用户token，注册Webhook时生成。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecurityToken *string `json:"security_token,omitempty"`
}

func (o PipelineTrigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineTrigger struct{}"
	}

	return strings.Join([]string{"PipelineTrigger", string(data)}, " ")
}
