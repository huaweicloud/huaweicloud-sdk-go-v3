package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BusinessTypePluginsQueryDto struct {

	// 用于区分插件为流水线可使用/模板可使用, 可选pipeline/template
	UseCondition *string `json:"use_condition,omitempty"`

	// 用于区分源的代码仓类型codehub/gitlab/github等
	InputRepoType *string `json:"input_repo_type,omitempty"`

	// 用于区分单源/多源的情况
	InputSourceType *string `json:"input_source_type,omitempty"`

	// 业务类型
	BusinessType *string `json:"business_type,omitempty"`

	// 名称
	RegexName *string `json:"regex_name,omitempty"`
}

func (o BusinessTypePluginsQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessTypePluginsQueryDto struct{}"
	}

	return strings.Join([]string{"BusinessTypePluginsQueryDto", string(data)}, " ")
}
