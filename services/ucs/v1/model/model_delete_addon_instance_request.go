package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAddonInstanceRequest Request Object
type DeleteAddonInstanceRequest struct {

	// 插件实例id
	Id string `json:"id"`

	// cluster_id 为空时，使用 addonInstance 的 cluster_id；不为空时，需要校验addonInstance 的 cluster_id 和 query 中的是否相同
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o DeleteAddonInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddonInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteAddonInstanceRequest", string(data)}, " ")
}
