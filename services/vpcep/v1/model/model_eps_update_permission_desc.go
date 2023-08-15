package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EpsUpdatePermissionDesc 更新终端节点服务白名单描述请求体。
type EpsUpdatePermissionDesc struct {

	// 终端节点服务白名单描述
	Description string `json:"description"`
}

func (o EpsUpdatePermissionDesc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsUpdatePermissionDesc struct{}"
	}

	return strings.Join([]string{"EpsUpdatePermissionDesc", string(data)}, " ")
}
