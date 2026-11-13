package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApsJobResourceInfo 任务结果对应的资源信息，根据任务类型区分： * `服务器` - 服务器信息 * `镜像` - 镜像信息
type ApsJobResourceInfo struct {

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 产物名称，或者被操作资源名称。
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ApsJobResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApsJobResourceInfo struct{}"
	}

	return strings.Join([]string{"ApsJobResourceInfo", string(data)}, " ")
}
