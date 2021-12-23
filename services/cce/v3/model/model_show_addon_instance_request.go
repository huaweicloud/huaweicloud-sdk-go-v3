package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAddonInstanceRequest struct {
	// 插件实例id

	Id string `json:"id"`
	// 集群 ID，获取方式请参见[[如何获取接口URI中参数](https://support.huaweicloud.com/api-cce/cce_02_0271.html)](tag:hws)[[如何获取接口URI中参数](https://support.huaweicloud.com/intl/zh-cn/api-cce/cce_02_0271.html)](tag:hws_hk)

	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ShowAddonInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAddonInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowAddonInstanceRequest", string(data)}, " ")
}
