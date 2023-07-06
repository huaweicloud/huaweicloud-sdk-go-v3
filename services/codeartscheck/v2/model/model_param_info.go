package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParamInfo 任务详情
type ParamInfo struct {

	// 仓库地址
	Url *string `json:"url,omitempty"`

	// 仓库分支
	Branch *string `json:"branch,omitempty"`

	// 仓库语言
	Language *string `json:"language,omitempty"`

	// 排除的目录
	ExcludeDir *string `json:"exclude_dir,omitempty"`

	// 编码格式
	Encode *string `json:"encode,omitempty"`

	// 编译配置信息
	CompileConfig *string `json:"compile_config,omitempty"`

	// g规则集名称
	RuleTemplate *string `json:"rule_template,omitempty"`
}

func (o ParamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamInfo struct{}"
	}

	return strings.Join([]string{"ParamInfo", string(data)}, " ")
}
