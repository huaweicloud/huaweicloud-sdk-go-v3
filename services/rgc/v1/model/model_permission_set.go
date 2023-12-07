package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PermissionSet IAM身份中心的权限集相关信息。
type PermissionSet struct {

	// 权限集ID。
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// 权限集名称。
	PermissionSetName *string `json:"permission_set_name,omitempty"`

	// 权限集描述。
	Description *string `json:"description,omitempty"`
}

func (o PermissionSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionSet struct{}"
	}

	return strings.Join([]string{"PermissionSet", string(data)}, " ")
}
