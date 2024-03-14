package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociatePermissionReqBody The request body of the AssociateResourceSharePermission operation.
type AssociatePermissionReqBody struct {

	// 指定特定的权限替换或绑定到与资源共享实例关联的现有资源类型。设置为\"true\"可将相同的资源类型的权限替换为当前权限。设置为\"false\"将权限绑定到当前资源类型。默认值为\"false\"。资源共享实例中的每个资源类型只能绑定一个权限。如果资源共享实例中已具有指定资源类型的权限，并且将\"replace\"设置为\"false\"，则操作返回错误。这有助于防止意外覆盖权限。
	Replace *bool `json:"replace,omitempty"`

	// 共享资源权限的ID。
	PermissionId string `json:"permission_id"`
}

func (o AssociatePermissionReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatePermissionReqBody struct{}"
	}

	return strings.Join([]string{"AssociatePermissionReqBody", string(data)}, " ")
}
