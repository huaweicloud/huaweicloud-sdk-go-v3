package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScopeRule struct {

	// 制品版本选择器，目前只支持长度为1
	TagSelectors []RuleSelector `json:"tag_selectors"`

	// 制品仓库选择器，目前只支持repository且长度为1
	ScopeSelectors map[string][]RuleSelector `json:"scope_selectors"`

	// repository选择方式。可选regular、selection，前端显示需要，api调用时可选regular
	RepoScopeMode string `json:"repo_scope_mode"`
}

func (o ScopeRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScopeRule struct{}"
	}

	return strings.Join([]string{"ScopeRule", string(data)}, " ")
}
