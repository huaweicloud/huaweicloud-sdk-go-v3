package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneAssociateGroupWithDomainPermissionRequest Request Object
type KeystoneAssociateGroupWithDomainPermissionRequest struct {

	// 用户组所属账号ID，获取方式请参见：[获取账号ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。
	DomainId string `json:"domain_id"`

	// 用户组ID，获取方式请参见：[获取用户组ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。
	GroupId string `json:"group_id"`

	// 权限ID，获取方式请参见：[获取权限名、权限ID](https://support.huaweicloud.com/api-iam/iam_10_0001.html)。
	RoleId string `json:"role_id"`
}

func (o KeystoneAssociateGroupWithDomainPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneAssociateGroupWithDomainPermissionRequest struct{}"
	}

	return strings.Join([]string{"KeystoneAssociateGroupWithDomainPermissionRequest", string(data)}, " ")
}
