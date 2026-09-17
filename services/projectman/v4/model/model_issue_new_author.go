package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueNewAuthor **参数解释：** 工作项负责人。 **取值范围：** 不涉及。
type IssueNewAuthor struct {

	// **参数解释：** 用户名称。 **取值范围：** 不涉及。
	FirstName *string `json:"firstName,omitempty"`

	// **参数解释：** 用户姓名。 **取值范围：** 不涉及。
	LastName *string `json:"lastName,omitempty"`

	// **参数解释：** 作者唯一标识。 **取值范围：** 不涉及。
	Identifier *string `json:"identifier,omitempty"`

	// **参数解释：** 用户头像id。 **取值范围：** 不涉及。
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释：** 用户昵称。 **取值范围：** 不涉及。
	AuthorNickName *string `json:"authorNickName,omitempty"`

	// **参数解释：** 带租户信息的用户名（租户名称_用户名）。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`
}

func (o IssueNewAuthor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueNewAuthor struct{}"
	}

	return strings.Join([]string{"IssueNewAuthor", string(data)}, " ")
}
