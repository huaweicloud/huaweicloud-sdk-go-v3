package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Image 云服务器镜像信息。
type Image struct {

	// 镜像ID。
	Id string `json:"id"`
}

func (o Image) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Image struct{}"
	}

	return strings.Join([]string{"Image", string(data)}, " ")
}
