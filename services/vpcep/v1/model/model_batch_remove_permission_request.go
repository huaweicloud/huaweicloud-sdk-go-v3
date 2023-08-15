package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemovePermissionRequest 批量删除终端节点服务白名单列表。
type BatchRemovePermissionRequest struct {
	Permissions []EpsRemovePermissionRequest `json:"permissions"`
}

func (o BatchRemovePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemovePermissionRequest struct{}"
	}

	return strings.Join([]string{"BatchRemovePermissionRequest", string(data)}, " ")
}
