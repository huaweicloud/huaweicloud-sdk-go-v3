package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPermissionRequest Request Object
type ShowPermissionRequest struct {

	// 共享资源权限的ID。
	PermissionId string `json:"permission_id"`
}

func (o ShowPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPermissionRequest struct{}"
	}

	return strings.Join([]string{"ShowPermissionRequest", string(data)}, " ")
}
