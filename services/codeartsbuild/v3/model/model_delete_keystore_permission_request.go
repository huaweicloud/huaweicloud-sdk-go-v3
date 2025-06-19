package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKeystorePermissionRequest Request Object
type DeleteKeystorePermissionRequest struct {

	// 文件管理权限ID
	PermissionId string `json:"permission_id"`
}

func (o DeleteKeystorePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeystorePermissionRequest struct{}"
	}

	return strings.Join([]string{"DeleteKeystorePermissionRequest", string(data)}, " ")
}
