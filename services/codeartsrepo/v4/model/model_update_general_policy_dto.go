package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateGeneralPolicyDto struct {

	// **参数解释：** 是否禁用fork。
	DisableFork *bool `json:"disable_fork,omitempty"`

	// **参数解释：** 分支名称正则表达式。 **取值范围：** 字符串长度不少于1，不超过1000。
	BranchNameRegex *string `json:"branch_name_regex,omitempty"`

	// **参数解释：** 标签名正则表达式。 **取值范围：** 字符串长度不少于1，不超过1000。
	TagNameRegex *string `json:"tag_name_regex,omitempty"`

	// **参数解释：** 生成合并请求预合并。
	GeneratePreMergeRef *bool `json:"generate_pre_merge_ref,omitempty"`
}

func (o UpdateGeneralPolicyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGeneralPolicyDto struct{}"
	}

	return strings.Join([]string{"UpdateGeneralPolicyDto", string(data)}, " ")
}
