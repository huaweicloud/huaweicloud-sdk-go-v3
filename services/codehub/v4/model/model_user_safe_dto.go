package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserSafeDto **参数解释：** 用户基础信息。
type UserSafeDto struct {

	// **参数解释：** 用户id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户名。
	Username *string `json:"username,omitempty"`
}

func (o UserSafeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserSafeDto struct{}"
	}

	return strings.Join([]string{"UserSafeDto", string(data)}, " ")
}
