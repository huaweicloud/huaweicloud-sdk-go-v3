package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemoveEndpointServicePermissionsRequestBody 批量删除终端节点服务白名单列表。
type BatchRemoveEndpointServicePermissionsRequestBody struct {

	// 终端节点服务白名单
	Permissions []EpsRemovePermissionRequest `json:"permissions"`
}

func (o BatchRemoveEndpointServicePermissionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveEndpointServicePermissionsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRemoveEndpointServicePermissionsRequestBody", string(data)}, " ")
}
