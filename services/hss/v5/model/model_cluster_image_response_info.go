package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterImageResponseInfo 集群防护镜像列表
type ClusterImageResponseInfo struct {

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本
	ImageVersion *string `json:"image_version,omitempty"`

	// ID
	Id *string `json:"id,omitempty"`
}

func (o ClusterImageResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterImageResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterImageResponseInfo", string(data)}, " ")
}
