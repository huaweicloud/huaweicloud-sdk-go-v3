package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppAclCreate struct {

	// 类型 -  PERMIT (白名单类型) -  DENY (黑名单类型)
	AppAclType string `json:"app_acl_type" xml:"app_acl_type"`

	// ACL策略值，支持IP、IP范围和CIDR方式。IP范围以英文中划线分隔。
	AppAclValues []string `json:"app_acl_values" xml:"app_acl_values"`
}

func (o AppAclCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppAclCreate struct{}"
	}

	return strings.Join([]string{"AppAclCreate", string(data)}, " ")
}
