package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectedRefsUserGroupBasicDto struct {

	// **参数解释：** 成员组ID。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 成员组名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ProtectedRefsUserGroupBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectedRefsUserGroupBasicDto struct{}"
	}

	return strings.Join([]string{"ProtectedRefsUserGroupBasicDto", string(data)}, " ")
}
