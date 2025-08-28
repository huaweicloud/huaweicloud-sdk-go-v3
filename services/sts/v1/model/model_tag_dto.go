package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagDto 标签，包含key、value两个字段。
type TagDto struct {

	// 标签键。
	Key string `json:"key"`

	// 标签值，取值可以为空字符串，不可以为null。
	Value string `json:"value"`
}

func (o TagDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagDto struct{}"
	}

	return strings.Join([]string{"TagDto", string(data)}, " ")
}
