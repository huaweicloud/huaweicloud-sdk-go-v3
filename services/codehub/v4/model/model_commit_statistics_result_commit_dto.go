package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommitStatisticsResultCommitDto struct {

	// **参数解释：** 作者名称。
	AuthorName *string `json:"author_name,omitempty"`

	// **参数解释：** 提交日期。
	Date *string `json:"date,omitempty"`

	// **参数解释：** 昵称。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 租户名。
	TenantName *string `json:"tenant_name,omitempty"`

	// **参数解释：** 用户名。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释：** 是否通过merge合入。 **取值范围：** - true，通过merge合入。 - false，非通过merge合入。
	IsMerge *bool `json:"is_merge,omitempty"`
}

func (o CommitStatisticsResultCommitDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitStatisticsResultCommitDto struct{}"
	}

	return strings.Join([]string{"CommitStatisticsResultCommitDto", string(data)}, " ")
}
