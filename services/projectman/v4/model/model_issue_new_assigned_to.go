package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueNewAssignedTo **参数解释：** 工作项责任人。 **取值范围：** 不涉及。
type IssueNewAssignedTo struct {

	// **参数解释：** 用户名。 **取值范围：** 不涉及。
	FirstName *string `json:"firstName,omitempty"`

	// **参数解释：** 用户姓名。 **取值范围：** 不涉及。
	LastName *string `json:"lastName,omitempty"`

	// **参数解释：** 用户32位uuid。 **取值范围：** 不涉及。
	Identifier *string `json:"identifier,omitempty"`

	// **参数解释：** 用户头像id。 **取值范围：** 不涉及。
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释：** 用户昵称。 **取值范围：** 不涉及。
	AssignedNickName *string `json:"assignedNickName,omitempty"`

	// **参数解释：** 用户名。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户数字id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`
}

func (o IssueNewAssignedTo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueNewAssignedTo struct{}"
	}

	return strings.Join([]string{"IssueNewAssignedTo", string(data)}, " ")
}
