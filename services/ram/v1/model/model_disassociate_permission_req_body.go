package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociatePermissionReqBody The request body of the DisassociateResourceSharePermission operation.
type DisassociatePermissionReqBody struct {

	// 共享资源权限的ID。
	PermissionId string `json:"permission_id"`
}

func (o DisassociatePermissionReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociatePermissionReqBody struct{}"
	}

	return strings.Join([]string{"DisassociatePermissionReqBody", string(data)}, " ")
}
