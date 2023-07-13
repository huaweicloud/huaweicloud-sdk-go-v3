package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppAclResponse Response Object
type UpdateAppAclResponse struct {

	// APP编号
	AppId *string `json:"app_id,omitempty"`

	// 类型 -  PERMIT (白名单类型) -  DENY (黑名单类型)
	AppAclType *string `json:"app_acl_type,omitempty"`

	// ACL策略值，支持IP、IP范围和CIDR方式。IP范围以英文中划线分隔。
	AppAclValues   *[]string `json:"app_acl_values,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateAppAclResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppAclResponse struct{}"
	}

	return strings.Join([]string{"UpdateAppAclResponse", string(data)}, " ")
}
