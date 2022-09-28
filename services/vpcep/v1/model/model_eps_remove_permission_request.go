package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除终端节点服务白名单请求体。
type EpsRemovePermissionRequest struct {

	// 终端节点服务白名单表主键ID
	Id string `json:"id"`
}

func (o EpsRemovePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsRemovePermissionRequest struct{}"
	}

	return strings.Join([]string{"EpsRemovePermissionRequest", string(data)}, " ")
}
