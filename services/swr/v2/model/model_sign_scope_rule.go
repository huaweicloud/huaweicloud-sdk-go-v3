package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SignScopeRule struct {

	// 制品版本选择器，目前只支持长度为1
	TagSelectors []SignRuleSelector `json:"tag_selectors"`

	// 制品仓库选择器，目前只支持repository且长度为1
	ScopeSelectors map[string][]SignRuleSelector `json:"scope_selectors"`

	// repository选择方式。可选regular、selection，前端显示需要，api调用时可选regular
	RepoScopeMode string `json:"repo_scope_mode"`
}

func (o SignScopeRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignScopeRule struct{}"
	}

	return strings.Join([]string{"SignScopeRule", string(data)}, " ")
}
