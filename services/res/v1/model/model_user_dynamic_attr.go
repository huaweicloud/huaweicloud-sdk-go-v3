package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type UserDynamicAttr struct {
	UserInterestedAttrs *Attribute `json:"user_interested_attrs,omitempty" xml:"user_interested_attrs"`

	UserInterestedActionType *Attribute `json:"user_interested_action_type,omitempty" xml:"user_interested_action_type"`
}

func (o UserDynamicAttr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserDynamicAttr struct{}"
	}

	return strings.Join([]string{"UserDynamicAttr", string(data)}, " ")
}
