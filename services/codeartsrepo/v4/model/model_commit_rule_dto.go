package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommitRuleDto struct {

	// **参数解释：** 主键ID。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库ID。 **取值范围：** 不涉及。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 提交信息匹配规则。
	CommitMessageRegex *string `json:"commit_message_regex,omitempty"`

	// **参数解释：** 提交信息负面匹配规则。
	CommitMessageNegativeRegex *string `json:"commit_message_negative_regex,omitempty"`

	// **参数解释：** 禁止提交的文件名称。
	ProhibitedFileNameRegex *string `json:"prohibited_file_name_regex,omitempty"`

	// **参数解释：** 提交人邮箱地址。
	AuthorEmailRegex *string `json:"author_email_regex,omitempty"`

	// **参数解释：** 单文件大小限制（MB）。
	MaxFileSize *int32 `json:"max_file_size,omitempty"`

	// **参数解释：** 允许的最大单文件大小（MB）。
	AllowedMaxFileSize *int32 `json:"allowed_max_file_size,omitempty"`

	// **参数解释：** 规则生效时间。
	EffectiveDate *string `json:"effective_date,omitempty"`

	// **参数解释：** 是否禁止新增二进制文件（对特权用户无效）。 **约束限制：** 不涉及。 **取值范围：** - true，禁止新增二进制文件。 - false，允许新增二进制文件。
	BinaryGateEnabled *bool `json:"binary_gate_enabled,omitempty"`

	// **参数解释：** 特权用户（特权用户可直接推送所有二进制文件入库）。
	PrivilegedUsers *[]RepositoryUserBasicDto `json:"privileged_users,omitempty"`

	// **参数解释：** 是否允许修改二进制文件（对特权用户无效）。 **约束限制：** 不涉及。 **取值范围：** - true，允许修改二进制文件。 - false，禁止修改二进制文件。
	AllowedModifyBinary *bool `json:"allowed_modify_binary,omitempty"`

	// **参数解释：** 二进制文件白名单（可直接入库的文件）。
	AllowedBinaryFileNameRegex *string `json:"allowed_binary_file_name_regex,omitempty"`

	// **参数解释：** 提交人。
	AuthorRegex *interface{} `json:"author_regex,omitempty"`

	// **参数解释：** 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 规则名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 分支规则。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	BranchName *string `json:"branch_name,omitempty"`

	// **参数解释：** 创建时间。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 跳过规则检测。 **约束限制：** 仅CR仓库支持此参数。
	SkipRuleCheck *bool `json:"skip_rule_check,omitempty"`

	// **参数解释：** 跳过规则检测失效时间， 例如: 2025-8-19。 **约束限制：** 仅CR仓库支持此参数。
	SkipRuleEndDate *string `json:"skip_rule_end_date,omitempty"`
}

func (o CommitRuleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitRuleDto struct{}"
	}

	return strings.Join([]string{"CommitRuleDto", string(data)}, " ")
}
