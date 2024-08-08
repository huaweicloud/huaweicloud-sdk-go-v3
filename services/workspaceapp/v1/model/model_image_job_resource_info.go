package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageJobResourceInfo 任务结果对应的资源信息，根据任务类型区分： * `镜像实例` - 镜像实例信息 * `镜像` - 镜像信息
type ImageJobResourceInfo struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 产物名称，或者被操作资源名称
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ImageJobResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageJobResourceInfo struct{}"
	}

	return strings.Join([]string{"ImageJobResourceInfo", string(data)}, " ")
}
