package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageTag 版本
type ImageTag struct {
}

func (o ImageTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTag struct{}"
	}

	return strings.Join([]string{"ImageTag", string(data)}, " ")
}
