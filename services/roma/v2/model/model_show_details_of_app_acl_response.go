package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDetailsOfAppAclResponse struct {

	// APP编号
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 类型 -  PERMIT (白名单类型) -  DENY (黑名单类型)
	AppAclType *string `json:"app_acl_type,omitempty" xml:"app_acl_type"`

	// ACL策略值，支持IP、IP范围和CIDR方式。IP范围以英文中划线分隔。
	AppAclValues   *[]string `json:"app_acl_values,omitempty" xml:"app_acl_values"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowDetailsOfAppAclResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAppAclResponse struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAppAclResponse", string(data)}, " ")
}
