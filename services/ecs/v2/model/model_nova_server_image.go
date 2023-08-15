package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaServerImage
type NovaServerImage struct {

	// 镜像ID。
	Id string `json:"id"`

	// 云服务器类型相关标记快捷链接信息。
	Links []NovaLink `json:"links"`
}

func (o NovaServerImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerImage struct{}"
	}

	return strings.Join([]string{"NovaServerImage", string(data)}, " ")
}
