package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageInfo image数据结构说明
type ImageInfo struct {

	// 镜像ID，格式为UUID。
	Id *string `json:"id,omitempty"`
}

func (o ImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageInfo struct{}"
	}

	return strings.Join([]string{"ImageInfo", string(data)}, " ")
}
