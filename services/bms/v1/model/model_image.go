package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// image字段数据结构说明
type Image struct {
	// 裸金属服务器镜像ID

	Id *string `json:"id,omitempty"`
	// 裸金属服务器镜像相关快捷链接信息

	Links *[]Links `json:"links,omitempty"`
}

func (o Image) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Image struct{}"
	}

	return strings.Join([]string{"Image", string(data)}, " ")
}
