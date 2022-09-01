package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务详情
type ParamInfo struct {

	// 仓库地址
	Url *string `json:"url,omitempty" xml:"url"`

	// 仓库分支
	Branch *string `json:"branch,omitempty" xml:"branch"`

	// 仓库语言
	Language *string `json:"language,omitempty" xml:"language"`

	// 排除的目录
	ExcludeDir *string `json:"exclude_dir,omitempty" xml:"exclude_dir"`

	// 编码格式
	Encode *string `json:"encode,omitempty" xml:"encode"`

	// 编译配置信息
	CompileConfig *string `json:"compile_config,omitempty" xml:"compile_config"`

	// g规则集名称
	RuleTemplate *string `json:"rule_template,omitempty" xml:"rule_template"`
}

func (o ParamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamInfo struct{}"
	}

	return strings.Join([]string{"ParamInfo", string(data)}, " ")
}
