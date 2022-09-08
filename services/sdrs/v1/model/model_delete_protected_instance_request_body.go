package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除保护实例请求体
type DeleteProtectedInstanceRequestBody struct {

	// 是否删除容灾站点服务器，默认值为false。
	DeleteTargetServer *bool `json:"delete_target_server,omitempty"`

	// 是否删除容灾站点弹性IP，默认值为false。
	DeleteTargetEip *bool `json:"delete_target_eip,omitempty"`
}

func (o DeleteProtectedInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectedInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteProtectedInstanceRequestBody", string(data)}, " ")
}
