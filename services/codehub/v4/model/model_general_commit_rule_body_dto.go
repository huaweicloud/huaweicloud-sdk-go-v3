package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GeneralCommitRuleBodyDto struct {

	// **参数解释：** 是否拒绝未Signed-off-by签名的提交。 **约束限制：** 不涉及。 **取值范围：** - true，拒绝未Signed-off-by签名的提交。 - false，允许未Signed-off-by签名的提交。 **默认取值：** 不涉及。
	RejectUnsignedCommits *bool `json:"reject_unsigned_commits,omitempty"`

	// **参数解释：** 是否拒绝未GPG签名的提交。 **约束限制：** 不涉及。 **取值范围：** - true，拒绝未GPG签名的提交。 - false，允许未GPG签名的提交。 **默认取值：** 不涉及。
	RejectNotSignedByGpg *bool `json:"reject_not_signed_by_gpg,omitempty"`

	// **参数解释：** 是否不允许删除Tags。 **约束限制：** 不涉及。 **取值范围：** - true，不允许删除Tags。 - false，允许删除Tags。 **默认取值：** 不涉及。
	DenyDeleteTag *bool `json:"deny_delete_tag,omitempty"`

	// **参数解释：** 是否阻止包含涉密文件的提交。 **约束限制：** 不涉及。 **取值范围：** - true，阻止包含涉密文件的提交。 - false，允许包含涉密文件的提交。 **默认取值：** 不涉及。
	PreventSecrets *bool `json:"prevent_secrets,omitempty"`

	// **参数解释：** 是否禁止强制推送。 **约束限制：** 不涉及。 **取值范围：** - true，禁止强制推送。 - false，允许强制推送。 **默认取值：** 不涉及。
	DenyForcePush *bool `json:"deny_force_push,omitempty"`
}

func (o GeneralCommitRuleBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralCommitRuleBodyDto struct{}"
	}

	return strings.Join([]string{"GeneralCommitRuleBodyDto", string(data)}, " ")
}
