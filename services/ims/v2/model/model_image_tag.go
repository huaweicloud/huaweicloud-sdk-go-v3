package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 镜像标签
type ImageTag struct {
	// 标签key值

	Key *string `json:"key,omitempty"`
	// 标签value值

	Value *string `json:"value,omitempty"`
}

func (o ImageTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTag struct{}"
	}

	return strings.Join([]string{"ImageTag", string(data)}, " ")
}
