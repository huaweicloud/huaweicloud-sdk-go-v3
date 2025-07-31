package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WhiteImageInfo 白名单镜像
type WhiteImageInfo struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 镜像名称
	ImageName string `json:"image_name"`

	// 镜像版本
	ImageVersion string `json:"image_version"`
}

func (o WhiteImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WhiteImageInfo struct{}"
	}

	return strings.Join([]string{"WhiteImageInfo", string(data)}, " ")
}
