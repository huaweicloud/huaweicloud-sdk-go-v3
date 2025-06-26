package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetentionRule struct {

	// 优先级，预留字段，默认为0
	Priority int32 `json:"priority"`

	// 是否关闭此条规则
	Disabled *bool `json:"disabled,omitempty"`

	// 预留字段，目前只支持retain
	Action string `json:"action"`

	// 规则模板类型，值为：latestPulledN，latestPushedK，nDaysSinceLastPush，nDaysSinceLastPull
	Template string `json:"template"`

	// 保留方式的参数，配套template使用，可参考请求示例
	Params map[string]interface{} `json:"params"`

	// 制品版本选择器，目前只支持长度为1
	TagSelectors []RetentionSelector `json:"tag_selectors"`

	// 制品仓库选择器，目前只支持repository且长度为1
	ScopeSelectors map[string][]RetentionSelector `json:"scope_selectors"`

	// repo选择模式，前端使用，可选regular、selection
	RepoScopeMode string `json:"repo_scope_mode"`
}

func (o RetentionRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetentionRule struct{}"
	}

	return strings.Join([]string{"RetentionRule", string(data)}, " ")
}
