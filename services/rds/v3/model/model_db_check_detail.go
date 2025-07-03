package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbCheckDetail 数据库检查结果
type DbCheckDetail struct {

	// 检查说明链接
	DocumentationLink *string `json:"documentation_link,omitempty"`

	// 检查项描述
	Description *string `json:"description,omitempty"`

	// 检查项id
	Id *string `json:"id,omitempty"`

	// 检查项标题
	Title *string `json:"title,omitempty"`

	// 检查状态
	Status *string `json:"status,omitempty"`

	// 各项检查项检测到的问题
	DetectedProblems *[]DetectedProblem `json:"detected_problems,omitempty"`
}

func (o DbCheckDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbCheckDetail struct{}"
	}

	return strings.Join([]string{"DbCheckDetail", string(data)}, " ")
}
