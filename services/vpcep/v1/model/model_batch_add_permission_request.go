package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量添加终端节点服务白名单列表。
type BatchAddPermissionRequest struct {

	// 终端节点服务白名单列表
	Permissions []EpsAddPermissionRequest `json:"permissions"`
}

func (o BatchAddPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddPermissionRequest struct{}"
	}

	return strings.Join([]string{"BatchAddPermissionRequest", string(data)}, " ")
}
