package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NovaServerImage struct {

	// 镜像ID。
	Id string `json:"id" xml:"id"`

	// 云服务器类型相关标记快捷链接信息。
	Links []NovaLink `json:"links" xml:"links"`
}

func (o NovaServerImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerImage struct{}"
	}

	return strings.Join([]string{"NovaServerImage", string(data)}, " ")
}
