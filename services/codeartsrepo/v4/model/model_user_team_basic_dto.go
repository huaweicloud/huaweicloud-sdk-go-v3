package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserTeamBasicDto struct {

	// **参数解释：** 用户组ID。 **取值范围：** 1-2147483647
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户组名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o UserTeamBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserTeamBasicDto struct{}"
	}

	return strings.Join([]string{"UserTeamBasicDto", string(data)}, " ")
}
