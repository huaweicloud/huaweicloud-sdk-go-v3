package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageName 镜像名称
type ImageName struct {
}

func (o ImageName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageName struct{}"
	}

	return strings.Join([]string{"ImageName", string(data)}, " ")
}
