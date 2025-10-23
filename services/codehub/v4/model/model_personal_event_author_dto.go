package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersonalEventAuthorDto struct {

	// **参数解释：** 用户id。 **约束限制：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名。 **约束限制：** 不涉及。
	Username *string `json:"username,omitempty"`
}

func (o PersonalEventAuthorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersonalEventAuthorDto struct{}"
	}

	return strings.Join([]string{"PersonalEventAuthorDto", string(data)}, " ")
}
