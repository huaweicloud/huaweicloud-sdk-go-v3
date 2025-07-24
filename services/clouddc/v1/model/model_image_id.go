package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageId 镜像ID，非必填，不传默认使用当前镜像ID
type ImageId struct {
}

func (o ImageId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageId struct{}"
	}

	return strings.Join([]string{"ImageId", string(data)}, " ")
}
