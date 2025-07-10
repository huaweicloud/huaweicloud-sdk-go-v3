package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdentityCenterGroup Identity Center的用户组信息。
type IdentityCenterGroup struct {

	// Identity Center的用户组ID。
	GroupId *string `json:"group_id,omitempty"`

	// 用户组名称。
	GroupName *string `json:"group_name,omitempty"`

	// Identity Center的用户组描述信息。
	Description *string `json:"description,omitempty"`
}

func (o IdentityCenterGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdentityCenterGroup struct{}"
	}

	return strings.Join([]string{"IdentityCenterGroup", string(data)}, " ")
}
