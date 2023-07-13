package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserDynamicAttr
type UserDynamicAttr struct {
	UserInterestedAttrs *Attribute `json:"user_interested_attrs,omitempty"`

	UserInterestedActionType *Attribute `json:"user_interested_action_type,omitempty"`
}

func (o UserDynamicAttr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserDynamicAttr struct{}"
	}

	return strings.Join([]string{"UserDynamicAttr", string(data)}, " ")
}
