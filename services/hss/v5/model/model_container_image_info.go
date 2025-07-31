package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerImageInfo 容器镜像信息
type ContainerImageInfo struct {

	// 镜像唯一标识(虚拟)
	ImageId *string `json:"image_id,omitempty"`

	// 镜像名
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本号
	ImageVersion *string `json:"image_version,omitempty"`

	// 镜像创建时间(本地保存时间)
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o ContainerImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerImageInfo struct{}"
	}

	return strings.Join([]string{"ContainerImageInfo", string(data)}, " ")
}
